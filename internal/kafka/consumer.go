package kafka

import (
	"context"
	"errors"
	"strings"

	"github.com/IBM/sarama"
)

// Consumer получатель-обработчик сообщений
type Consumer interface {
	RunConsume(ctx context.Context, topic string) error
	GroupHandler() *GroupHandler
	Close() error
}

type consumer struct {
	consumer sarama.ConsumerGroup
	handler  *GroupHandler
}

// NewConsumer новый экземпляр
func NewConsumer(brokers []string, groupID string, config *sarama.Config) (Consumer, error) {
	cons, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &consumer{consumer: cons, handler: NewGroupHandler()}, nil
}

func (c *consumer) GroupHandler() *GroupHandler {
	return c.handler
}

func (c *consumer) RunConsume(ctx context.Context, topic string) error {
	for {
		err := c.consumer.Consume(ctx, strings.Split(topic, ","), c.handler)
		if err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}

			if errors.Is(err, ErrRebalancingGroup) {
				continue
			}

			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *consumer) Close() error {
	return c.consumer.Close()
}
