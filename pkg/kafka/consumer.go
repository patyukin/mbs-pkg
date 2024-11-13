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

func (c *Client) ProcessMessages(ctx context.Context, processFunc func(*kgo.Record) error) error {
	for {
		fetches := c.client.PollFetches(ctx)
		if fetches.IsClientClosed() {
			return nil
		}

		var errs []error
		fetches.EachError(func(t string, p int32, err error) {
			errs = append(errs, fmt.Errorf("failed to fetch %s partition %d: %w", t, p, err))
		})

		if len(errs) > 0 {
			return errors.Join(errs...)
		}

		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			for _, record := range p.Records {
				go c.processRecord(ctx, record, processFunc)
			}
		})
	}
}

func (c *Client) processRecord(ctx context.Context, record *kgo.Record, processFunc func(*kgo.Record) error) {
	var attempt int
	backoff := time.Millisecond * 100
	maxBackoff := time.Second * 2

	for attempt = 1; attempt <= maxRetries; attempt++ {
		err := processFunc(record)
		if err == nil {
			log.Printf("Сообщение из топика %s успешно обработано\n", record.Topic)
			return
		}

		log.Printf("Попытка обработки %d не удалась: %v", attempt, err)

		if attempt < maxRetries {
			time.Sleep(backoff)
			backoff = time.Duration(math.Min(float64(maxBackoff), float64(backoff)*2))
		} else {
			log.Error().Msgf("Сообщение отправлено в dead-letter топик %s после %d неудачных попыток\n", DeadLetterTopic, maxRetries)
			c.sendToDeadLetter(ctx, record)
		}
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
			log.Error().Msgf("Ошибка при отправке в dead-letter топик: %v", err)
		} else {
			log.Info().Msgf("Сообщение успешно отправлено в dead-letter топик %s\n", DeadLetterTopic)
		}
	})
}
