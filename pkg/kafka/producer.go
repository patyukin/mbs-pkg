package kafka

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"math"
	"time"
)

func (c *Client) PublishMessageWithRetry(ctx context.Context, topic string, key, value []byte) error {
	backoff := time.Millisecond * 100
	record := &kgo.Record{Topic: topic, Key: key, Value: value}

	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		resultCh := make(chan error, 1)

		c.client.Produce(ctx, record, func(_ *kgo.Record, err error) {
			resultCh <- err
		})

		err := <-resultCh
		if err == nil {
			log.Info().Msgf("Message %s sent to topic %s", string(value), topic)
			return nil
		}

		lastErr = err
		log.Error().Msgf("Attempt %d failed for message %s to topic %s: %v", attempt, string(value), topic, err)

		if attempt == maxRetries {
			break
		}

		time.Sleep(backoff)
		backoff = time.Duration(math.Min(float64(maxBackoff), float64(backoff)*2))
	}

	log.Error().Msgf("Message %s not sent to topic %s after %d attempts", string(value), topic, maxRetries)

	return lastErr
}

func (c *Client) PublishLogReport(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, LogReportTopic, nil, value)
}

func (c *Client) PublishTransaction(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, TransactionTopic, nil, value)
}

func (c *Client) PublishCreditCreated(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, CreditCreatedTopic, nil, value)
}

func (c *Client) PublishCreditPayments(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, CreditPaymentsTopic, nil, value)
}

func (c *Client) PublishPaymentRequest(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, PaymentRequestTopic, nil, value)
}
