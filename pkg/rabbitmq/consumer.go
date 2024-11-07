package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func (r *RabbitMQ) Consume(ctx context.Context, queue string, processMessage HandlerFunction) error {
	msgs, err := r.channel.Consume(
		queue,
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

	for msg := range msgs {
		go func(d amqp.Delivery) {
			if err = processMessage(ctx, d); err != nil {
				if err = d.Nack(false, false); err != nil {
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

	return nil
}
