package consumer

import (
	"io"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
	writer   io.Writer
}

func New(addrs []string, w io.Writer) (*Consumer, error) {
	cfg := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(addrs, cfg)
	if err != nil {
		return nil, err
	}

	return &Consumer{consumer, w}, nil
}

func (c *Consumer) Consume(topic string, partition int32, offset int64, quit chan int) error {
	pc, err := c.consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case data := <-pc.Messages():
				// TODO: interested in key?
				_, err = c.writer.Write(data.Value)
				if err != nil {
					return
				}
			case <-quit:
				return
			}
		}
	}()

	return nil
}
