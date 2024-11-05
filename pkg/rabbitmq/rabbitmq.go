package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	authSignUpConfirmCodeRouteKey = "auth_sign_up_confirm_code_route_key"
	authSignUpConfirmCodeQueue    = "auth_sign_up_confirm_code_queue"
	authSignUpRouteKey            = "auth_sign_up_route_key"
	authSignInQueue               = "auth_sign_in_queue"
)

type HandlerFunction func(amqp.Delivery) error

type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
}

func New(url, exchange string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	err = channel.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return &RabbitMQ{
		conn:     conn,
		channel:  channel,
		exchange: exchange,
	}, nil
}

func (r *RabbitMQ) BindQueueToExchange(exchangeName, queueName string, routingKeys []string) error {
	q, err := r.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-message-ttl": int64(2592000000), // 30 дней = 30 * 24 * 60 * 60 * 1000 = 2592000000 миллисекунд
		},
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	for _, key := range routingKeys {
		err = r.channel.QueueBind(
			q.Name,
			key,
			exchangeName,
			false,
			nil,
		)
		if err != nil {
			return fmt.Errorf("failed to bind queue to exchange: %w", err)
		}
	}

	return nil
}

func (r *RabbitMQ) QueuesDeclare(queues []string) error {
	for _, queue := range queues {
		_, err := r.channel.QueueDeclare(
			queue,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return fmt.Errorf("failed to declare queue: %w", err)
		}
	}

	return nil

}

func (r *RabbitMQ) Close() error {
	return r.conn.Close()
}

func (r *RabbitMQ) CloseChannel() error {
	return r.channel.Close()
}

func (r *RabbitMQ) AuthSignInCodeConsumer(processMessage HandlerFunction) {
	msgs, err := r.channel.Consume(
		authSignInQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Msgf("failed to consume: %v", err)
		return
	}

	for msg := range msgs {
		go func(d amqp.Delivery) {
			if err = processMessage(d); err != nil {
				if err = d.Nack(false, true); err != nil {
					log.Error().Msgf("failed to nack message: %v", err)
					return
				}

				log.Error().Msgf("failed to process message: %v", err)
				return
			}

			if err = d.Ack(false); err != nil {
				log.Error().Msgf("failed to ack message: %v", err)
				return
			}
		}(msg)
	}
}

func (r *RabbitMQ) AuthSignUpConfirmCodeConsumer(processMessage HandlerFunction) {
	msgs, err := r.channel.Consume(
		authSignUpConfirmCodeQueue,
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
			if err = processMessage(d); err != nil {
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

func (r *RabbitMQ) publishMessage(ctx context.Context, routeKey string, body []byte) error {
	err := r.channel.PublishWithContext(
		ctx,
		r.exchange,
		routeKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Info().Msgf("Message sent to route key %s via exchange %s: %s", routeKey, r.exchange, string(body))

	return nil
}

func (r *RabbitMQ) PublishAuthSignUpMessage(ctx context.Context, body []byte) error {
	return r.publishMessage(ctx, authSignUpRouteKey, body)
}

func (r *RabbitMQ) PublishAuthSignUpConfirmCodeMessage(ctx context.Context, body []byte) error {
	return r.publishMessage(ctx, authSignUpConfirmCodeRouteKey, body)
}
