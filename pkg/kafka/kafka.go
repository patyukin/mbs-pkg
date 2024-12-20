package kafka

import (
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Client struct {
	client *kgo.Client
}

func NewConsumer(brokers []string, consumerGroup string, topics []string) (*Client, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(consumerGroup),
		kgo.ConsumeTopics(topics...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed create consumer, err: %v", err)
	}

	return &Client{client: cl}, err
}

func NewProducer(brokers []string) (*Client, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ProducerBatchCompression(kgo.GzipCompression()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed create producer, err: %v", err)
	}

	return &Client{client: cl}, err
}

func (c *Client) Close() {
	c.client.Close()
}
