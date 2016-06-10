package producer

import (
	"bufio"
	"errors"
	"io"
	"os"

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
	buf := make([]byte, 1024)
	var err error
	for err == nil {
		_, err = p.reader.Read(buf)
		if err != nil {
			return nil
		}
		// send to producer here
	}
	return nil
}

type LineReader struct {
	scanner *bufio.Scanner
}

func NewLineReader(f string) (*LineReader, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	// TODO: handle closing of file

	scanner := bufio.NewScanner(file)
	return &LineReader{scanner}, nil
}

func (r *LineReader) Read(p []byte) (int, error) {
	if !r.scanner.Scan() {
		return 0, errors.New("file read done")
	}
	str := r.scanner.Text()
	copy(p, []byte(str))
	return len(str), nil
}
