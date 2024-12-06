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
		return fmt.Errorf("failed to start consuming from queue '%s': %w", queue, err)
	}

	log.Info().Msgf("Started consuming from queue '%s'", queue)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info().Msgf("Stopping consumption from queue '%s' due to context cancellation", queue)
				return
			case msg, ok := <-msgs:
				if !ok {
					log.Warn().Msgf("Message channel closed for queue '%s'", queue)
					return
				}

				go r.handleMessage(ctx, queue, msg, processMessage)
			}
		}
	}()

	return nil
}

func (r *RabbitMQ) handleMessage(ctx context.Context, queue string, msg amqp.Delivery, processMessage HandlerFunction) {
	if err := processMessage(ctx, msg); err != nil {
		if nackErr := msg.Nack(false, false); nackErr != nil {
			log.Error().
				Err(nackErr).
				Msgf("Failed to nack message from queue '%s'", queue)
			return
		}

		log.Error().
			Err(err).
			Msgf("Failed to process message from queue '%s'", queue)
		return
	}

	if ackErr := msg.Ack(false); ackErr != nil {
		log.Error().
			Err(ackErr).
			Msgf("Failed to ack message from queue '%s'", queue)
		return
	}

	log.Debug().
		Str("queue", queue).
		Str("message", string(msg.Body)).
		Msg("Message processed successfully")
}
