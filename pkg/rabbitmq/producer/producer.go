package producer

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
}

func New(url, exchange string) (*Producer, error) {
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

	return &Producer{
		conn:     conn,
		channel:  channel,
		exchange: exchange,
	}, nil
}

func (p *Producer) BindQueueToExchange(exchangeName, queueName string, routingKeys []string) error {
	for _, key := range routingKeys {
		err := p.channel.QueueBind(
			queueName,
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

func (p *Producer) PublishMessage(ctx context.Context, routeKey string, body []byte) error {
	err := p.channel.PublishWithContext(
		ctx,
		p.exchange,
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

	log.Info().Msgf("Message sent to route key %s via exchange %s: %s", routeKey, p.exchange, string(body))

	return nil
}
