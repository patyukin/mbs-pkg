package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
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

	err = channel.ExchangeDeclare(
		DeadLetterExchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a DLX exchange: %w", err)
	}

	_, err = channel.QueueDeclare(
		DeadLetterQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a DLQ: %w", err)
	}

	err = channel.QueueBind(
		DeadLetterQueue,
		DeadLetterQueue,
		DeadLetterExchange,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind DLQ to DLX: %w", err)
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
			"x-dead-letter-exchange":    DeadLetterExchange,
			"x-dead-letter-routing-key": DeadLetterQueue,
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
