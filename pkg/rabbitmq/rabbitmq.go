package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	NotifySignUpConfirmCodeRouteKey  = "notify_sign_up_confirm_code_route_key"
	AuthSignUpConfirmMessageRouteKey = "auth_sign_up_confirm_message_route_key"
	AuthSignInCodeRouteKey           = "auth_sign_in_code_route_key"
	AuthNotifyQueue                  = "auth_notify_queue"
	NotifyAuthQueue                  = "notify_auth_queue"
	dlqRouteKey                      = "dead_letter_queue_route_key"
)

type HandlerFunction func(ctx context.Context, d amqp.Delivery) error

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

func (r *RabbitMQ) BindQueueToExchange(exchangeName, queueName string, routingKeys []string, workerCount int) error {
	q, err := r.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "",
			"x-dead-letter-routing-key": dlqRouteKey,
			"x-message-ttl":             int64(2592000000), // 30 дней = 30 * 24 * 60 * 60 * 1000 = 2592000000 миллисекунд
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

	if workerCount > 0 {
		err = r.channel.Qos(
			workerCount,
			0,
			false,
		)
		if err != nil {
			return fmt.Errorf("failed to set QoS: %w", err)
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
