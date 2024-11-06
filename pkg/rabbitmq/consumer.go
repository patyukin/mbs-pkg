package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"sync"
)

// AuthNotifyConsumer - consumer for auth_notify_queue
// Обработка сообщений из очереди auth_notify_queue
func (r *RabbitMQ) AuthNotifyConsumer(ctx context.Context, processMessage HandlerFunction, workerCount int) error {
	msgs, err := r.channel.Consume(
		AuthNotifyQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume in auth_notify_queue: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(workerCount)

	// run workers
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Error().Msgf("worker %d: context is done", workerID)
					return
				case msg, ok := <-msgs:
					if !ok {
						log.Error().Msgf("worker %d: channel is closed", workerID)
						return
					}

					if err = processMessage(ctx, msg); err != nil {
						if nackErr := msg.Nack(false, true); nackErr != nil {
							log.Printf("worker %d: failed to nack message: %v", workerID, nackErr)
							continue
						}

						log.Printf("worker %d: failed to process message: %v", workerID, err)
						continue
					}

					if ackErr := msg.Ack(false); ackErr != nil {
						log.Printf("worker %d: failed to ack message: %v", workerID, ackErr)
					}
				}
			}
		}(i + 1)
	}

	wg.Wait()
	log.Info().Msg("All workers finished")

	return nil
}

func (r *RabbitMQ) NotifyAuthConsumer(ctx context.Context, processMessage HandlerFunction) {
	msgs, err := r.channel.Consume(
		NotifyAuthQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Msgf("failed to consume in auth_sign_up_confirm_code_queue: %v", err)
		return
	}

	for msg := range msgs {
		go func(d amqp.Delivery) {
			if err = processMessage(ctx, d); err != nil {
				if err = d.Nack(false, true); err != nil {
					log.Error().Msgf("failed to nack message in auth_sign_up_confirm_code_queue: %v", err)
					return
				}

				log.Error().Msgf("failed to process message in auth_sign_up_confirm_code_queue: %v", err)
				return
			}

			if err = d.Ack(false); err != nil {
				log.Error().Msgf("failed to ack message in auth_sign_up_confirm_code_queue: %v", err)
				return
			}

			log.Debug().Msgf("Message processed in auth_sign_up_confirm_code_queue: %s", string(d.Body))
		}(msg)
	}
}
