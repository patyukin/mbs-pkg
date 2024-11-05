package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"time"
)

func (r *RabbitMQ) publishMessage(ctx context.Context, routeKey string, body []byte) error {
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
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Info().Msgf("Message sent to route key %s via exchange %s: %s", routeKey, r.exchange, string(body))

	return nil
}

// PublishSignUpCodeRouteKeyMessage
// Отправка кода подтверждения регистрации из телеграм в RabbitMQ
func (r *RabbitMQ) PublishSignUpCodeRouteKeyMessage(ctx context.Context, body []byte) error {
	return r.publishMessage(ctx, SignUpConfirmCodeRouteKey, body)
}

// PublishAuthSignUpMessage
// Отправка информации подтверждения регистрации из authService в RabbitMQ
func (r *RabbitMQ) PublishAuthSignUpMessage(ctx context.Context, body []byte) error {
	return r.publishMessage(ctx, SignUpConfirmMessageRouteKey, body)
}

// PublishSignInCodeRouteKeyMessage
// Отправка кода подтверждения входа из authService в RabbitMQ
func (r *RabbitMQ) PublishSignInCodeRouteKeyMessage(ctx context.Context, body []byte) error {
	return r.publishMessage(ctx, SignInCodeRouteKey, body)
}
