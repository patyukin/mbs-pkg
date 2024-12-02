package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"math"
	"time"
)

func (c *Client) ProcessMessages(ctx context.Context, processFunc func(ctx context.Context, record *kgo.Record) error) error {
	for {
		fetches := c.client.PollFetches(ctx)
		if fetches.IsClientClosed() {
			log.Debug().Msg("Client is closed")
			return nil
		}

		var errs []error
		fetches.EachError(func(t string, p int32, err error) {
			errs = append(errs, fmt.Errorf("failed to fetch %s partition %d: %w", t, p, err))
		})

		log.Debug().Msgf("Received %d fetch errors", len(errs))

		if len(errs) > 0 {
			return errors.Join(errs...)
		}

		log.Debug().Msgf("Received %d fetch responses", fetches.NumRecords())

		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			for _, record := range p.Records {
				go c.processRecord(ctx, record, processFunc)
			}
		})
	}
}

func (c *Client) processRecord(ctx context.Context, record *kgo.Record, processFunc func(context.Context, *kgo.Record) error) {
	var attempt int
	backoff := time.Millisecond * 100

	for attempt = 1; attempt <= maxRetries; attempt++ {
		err := processFunc(ctx, record)
		if err == nil {
			log.Info().Msgf("The message was successfully sent to the topic %s: %s", record.Topic, string(record.Value))
			return
		}

		log.Error().Msgf("Attempt %d failed when sending a message to the topic %s: %v", attempt, record.Topic, err)

		if attempt >= maxRetries {
			log.Error().Msgf("Failed to send message to topic %s after %d attempts: %v", record.Topic, maxRetries, err)
			c.sendToDeadLetter(ctx, record)
			return
		}

		time.Sleep(backoff)
		backoff = time.Duration(math.Min(float64(maxBackoff), float64(backoff)*2))
	}
}

// sendToDeadLetter отправляет сообщение в dead-letter топик
func (c *Client) sendToDeadLetter(ctx context.Context, record *kgo.Record) {
	deadLetterRecord := &kgo.Record{
		Topic:   DeadLetterTopic,
		Key:     record.Key,
		Value:   record.Value,
		Headers: record.Headers,
	}

	c.client.Produce(ctx, deadLetterRecord, func(_ *kgo.Record, err error) {
		if err != nil {
			log.Error().Msgf("Failed to send message to dead-letter topic: %v", err)
			return
		}

		log.Info().Msgf("Successfully sent message to dead-letter topic: %s", DeadLetterTopic)
	})
}
