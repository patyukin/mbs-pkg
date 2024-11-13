package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"time"
)

// PublishDQLMessage
// Отправка сообщение в DQL в RabbitMQ
func (r *RabbitMQ) PublishDQLMessage(ctx context.Context, body []byte) error {
	err := r.channel.PublishWithContext(
		ctx,
		DeadLetterExchange,
		DeadLetterQueue,
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

	log.Info().Msgf("Message sent to route key %s via exchange %s: %s", DeadLetterQueue, r.exchange, string(body))

	return nil
}

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

// PublishNotifySignUpConfirmCode
// Отправка кода подтверждения регистрации из телеграм в RabbitMQ
func (r *RabbitMQ) PublishNotifySignUpConfirmCode(ctx context.Context, body []byte, headers amqp.Table) error {
	return r.publishMessage(ctx, NotifySignUpConfirmCodeRouteKey, body, headers)
}

// PublishAuthSignUpResultMessage
// Отправка информации подтверждения регистрации из authService в RabbitMQ
func (r *RabbitMQ) PublishAuthSignUpResultMessage(ctx context.Context, body []byte, headers amqp.Table) error {
	return r.publishMessage(ctx, AuthSignUpResultMessageRouteKey, body, headers)
}

// PublishAuthSignInCode
// Отправка кода подтверждения входа из authService в RabbitMQ
func (r *RabbitMQ) PublishAuthSignInCode(ctx context.Context, body []byte, headers amqp.Table) error {
	return r.publishMessage(ctx, AuthSignInConfirmCodeRouteKey, body, headers)
}

// PublishLogReport
// Отправка информации о логах сервиса в RabbitMQ
func (r *RabbitMQ) PublishLogReport(ctx context.Context, body []byte, headers amqp.Table) error {
	return r.publishMessage(ctx, LoggerReportRouteKey, body, headers)
}
