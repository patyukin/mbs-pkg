package kafka

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"math"
	"sync"
	"time"
)

// ProcessMessages - process messages from kafka topic
func (c *Client) ProcessMessages(ctx context.Context, processFunc func(ctx context.Context, record *kgo.Record) error) error {
	sem := make(chan struct{}, maxGoroutines)
	wg := &sync.WaitGroup{}
	errorCh := make(chan error, maxGoroutines)

	go func() {
		for err := range errorCh {
			log.Error().Msgf("Error processing message: %v", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			close(errorCh)
			return ctx.Err()
		default:
			fetches := c.client.PollFetches(ctx)
			if fetches.IsClientClosed() {
				wg.Wait()
				close(errorCh)
				return nil
			}

			var fetchErrors []error
			fetches.EachError(func(t string, p int32, err error) {
				fetchErrors = append(fetchErrors, fmt.Errorf("failed to fetch from topic %s partition %d: %w", t, p, err))
			})

			if len(fetchErrors) > 0 {
				for _, err := range fetchErrors {
					log.Error().Msgf("Error fetching messages: %v", err)
				}
			}

			fetches.EachPartition(func(p kgo.FetchTopicPartition) {
				for _, record := range p.Records {
					select {
					case <-ctx.Done():
						return
					case sem <- struct{}{}:
						wg.Add(1)

						go func(rec *kgo.Record) {
							defer wg.Done()
							defer func() { <-sem }()

							if err := c.processRecord(ctx, rec, processFunc); err != nil {
								select {
								case errorCh <- err:
								default:
									log.Error().Msgf("many errors in channel, dropping error: %v", err)
								}
							}
						}(record)
					}
				}
			})
		}
	}
}

// processRecord - process one record
func (c *Client) processRecord(ctx context.Context, record *kgo.Record, processFunc func(context.Context, *kgo.Record) error) error {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("panic in processRecord: %v", r)
			if err := c.sendToDeadLetter(ctx, record); err != nil {
				log.Error().Msgf("error sending to dead letter: %v", err)
			}
		}
	}()

	select {
	case <-ctx.Done():
		log.Error().Msgf("Context canceled, record processing stopped, err: %v", ctx.Err())
		return ctx.Err()
	default:
	}

	err := processFunc(ctx, record)
	if err == nil {
		log.Info().Msgf("Successfully processed message from topic %s", record.Topic)
		return nil
	}

	log.Info().Msgf("processing failed, err: %v, Sending a message to dead-letter topic %s after %d failed attempts", err, DeadLetterTopic, maxRetries)
	if err = c.sendToDeadLetter(ctx, record); err != nil {
		return fmt.Errorf("failed to send a message to the dead-letter topic: %w", err)
	}

	return nil
}

// sendToDeadLetter - send message to dead letter
func (c *Client) sendToDeadLetter(ctx context.Context, record *kgo.Record) error {
	deadLetterRecord := &kgo.Record{
		Topic:   DeadLetterTopic,
		Key:     record.Key,
		Value:   record.Value,
		Headers: record.Headers,
	}

	backoff := time.Millisecond * 100

	for attempt := 1; attempt <= maxRetries; attempt++ {
		errCh := make(chan error, 1)
		done := make(chan struct{})

		c.client.Produce(ctx, deadLetterRecord, func(_ *kgo.Record, err error) {
			if err != nil {
				log.Error().Msgf("error sending to dead letter: %v", err)
				errCh <- err
			} else {
				log.Info().Msgf("success sending to dead letter: %s", string(record.Value))
				errCh <- nil
			}

			close(done)
		})

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-done:
			err := <-errCh
			if err == nil {
				return nil
			}

			log.Info().Msgf("Attempt %d to send topic to dead-letter failed: %v", attempt, err)
			if attempt >= maxRetries {
				log.Error().Msgf("Failed to send a message to the dead-letter topic after %d attempts", err)
				return err
			}

			select {
			case <-time.After(backoff):
				backoff = time.Duration(math.Min(float64(maxBackoff), float64(backoff)*2))
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	return fmt.Errorf("sendToDeadLetter: The maximum number of attempts has been exceeded")
}
