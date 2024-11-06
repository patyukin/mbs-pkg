package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"time"
)

func (r *RabbitMQ) publishMessage(ctx context.Context, routeKey string, body []byte, headers amqp.Table) error {
	err := r.channel.PublishWithContext(
		ctx,
		r.exchange,
		routeKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			Timestamp:    time.Now(),
			DeliveryMode: amqp.Persistent,
			Headers:      headers,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Info().Msgf("Message sent to route key %s via exchange %s: %s", routeKey, r.exchange, string(body))

	return nil
}

// PublishDQLMessage
// Отправка сообщение в DQL в RabbitMQ
func (r *RabbitMQ) PublishDQLMessage(ctx context.Context, body []byte) error {
	err := r.channel.PublishWithContext(
		ctx,
		"",
		DeadLetterQueueRouteKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			Timestamp:    time.Now(),
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Info().Msgf("Message sent to route key %s via exchange %s: %s", DeadLetterQueueRouteKey, r.exchange, string(body))

	return nil
}

// PublishSignUpCodeRouteKeyMessage
// Отправка кода подтверждения регистрации из телеграм в RabbitMQ
func (r *RabbitMQ) PublishSignUpCodeRouteKeyMessage(ctx context.Context, body []byte, headers amqp.Table) error {
	return r.publishMessage(ctx, NotifySignUpConfirmCodeRouteKey, body, headers)
}

// PublishAuthSignUpMessage
// Отправка информации подтверждения регистрации из authService в RabbitMQ
func (r *RabbitMQ) PublishAuthSignUpMessage(ctx context.Context, body []byte, headers amqp.Table) error {
	return r.publishMessage(ctx, AuthSignUpConfirmMessageRouteKey, body, headers)
}

// PublishSignInCodeRouteKeyMessage
// Отправка кода подтверждения входа из authService в RabbitMQ
func (r *RabbitMQ) PublishSignInCodeRouteKeyMessage(ctx context.Context, body []byte, headers amqp.Table) error {
	return r.publishMessage(ctx, AuthSignInCodeRouteKey, body, headers)
}
