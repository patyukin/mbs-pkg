package kafka

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"math"
	"time"
)

// PublishMessageWithRetry - sends a message to the specified Kafka topic with repeated attempts if it fails.
func (c *Client) PublishMessageWithRetry(ctx context.Context, topic string, key, value []byte) error {
	backoff := initialBackoff
	record := &kgo.Record{Topic: topic, Key: key, Value: value}

	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := c.produceMessage(ctx, record)
		if err == nil {
			log.Info().Msgf("The message was successfully sent to the topic %s: %s", topic, string(value))
			return nil
		}

		lastErr = err
		log.Info().Msgf("Attempt %d failed when sending a message to the topic %s: %v", attempt, topic, err)

		if attempt < maxRetries {
			select {
			case <-time.After(backoff):
				backoff = time.Duration(math.Min(float64(maxBackoff), float64(backoff)*2))
			case <-ctx.Done():
				log.Error().Msgf("Context cancellation while waiting before trying %d for topic %s: %v", attempt+1, topic, ctx.Err())
				return ctx.Err()
			}
		}
	}

	log.Info().Msgf("Failed to send message to topic %s after %d attempts: %v", topic, maxRetries, lastErr)
	return lastErr
}

// produceMessage synchronously sends a message to Kafka and returns an error if it has occurred.
func (c *Client) produceMessage(ctx context.Context, record *kgo.Record) error {
	resultCh := make(chan error, 1)

	c.client.Produce(ctx, record, func(_ *kgo.Record, err error) {
		resultCh <- err
	})

	select {
	case err := <-resultCh:
		return fmt.Errorf("failed to produce message: %w", err)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *Client) PublishLogReport(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, LogReportTopic, nil, value)
}

func (c *Client) PublishPaymentRequest(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, PaymentRequestTopic, nil, value)
}

func (c *Client) PublishCreditPaymentSolution(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, CreditPaymentsSolutionTopic, nil, value)
}

func (c *Client) PublishTransactionReport(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, TransactionReportTopic, nil, value)
}

func (c *Client) PublishCreditCreated(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, CreditCreatedTopic, nil, value)
}

func (c *Client) PublishCreditPayments(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, CreditPaymentsTopic, nil, value)
}

func (c *Client) PublishCreditPaymentsSolution(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, CreditPaymentsSolutionTopic, nil, value)
}

func (c *Client) PublishRegistrationSolution(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, RegistrationSolutionTopic, nil, value)
}
