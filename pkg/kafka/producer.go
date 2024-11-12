package kafka

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"math"
	"sync"
	"time"
)

func (c *Client) PublishMessageWithRetry(ctx context.Context, topic string, key, value []byte) error {
	backoff := time.Millisecond * 100
	record := &kgo.Record{Topic: topic, Key: key, Value: value}

	var attempt int
	var wg sync.WaitGroup
	wg.Add(1)

	resultCh := make(chan error, 1)

	go func() {
		defer wg.Done()
		for attempt = 1; attempt <= maxRetries; attempt++ {
			c.client.Produce(ctx, record, func(_ *kgo.Record, err error) {
				resultCh <- err
			})

			err := <-resultCh
			if err == nil {
				log.Printf("Сообщение успешно отправлено в топик %s\n", topic)
				return
			}

			log.Printf("Попытка %d не удалась: %v", attempt, err)

			if attempt < maxRetries {
				log.Printf("Повторная попытка через %v...\n", backoff)
				time.Sleep(backoff)
				backoff = time.Duration(math.Min(float64(maxBackoff), float64(backoff)*2))
			} else {
				log.Printf("Все %d попыток отправки сообщения в топик %s исчерпаны.\n", maxRetries, topic)
			}
		}
	}()

	wg.Wait()

	if attempt > maxRetries {
		return <-resultCh
	}

	return nil
}

func (c *Client) PublishLogReport(ctx context.Context, value []byte) error {
	return c.PublishMessageWithRetry(ctx, LogReportTopic, nil, value)
}
