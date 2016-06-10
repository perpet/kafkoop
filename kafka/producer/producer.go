package producer

import (
	"io"

	"github.com/Shopify/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
	reader   io.Reader
}

func New(addrs []string, r io.Reader) (*Producer, error) {
	cfg := sarama.NewConfig()
	p, err := sarama.NewSyncProducer(addrs, cfg)
	if err != nil {
		return nil, err
	}

	return &Producer{p, r}, nil
}

func (p *Producer) Produce() error {
	// read from p.reader and output to p.producer
	return nil
}
