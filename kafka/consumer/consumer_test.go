package consumer

import (
	"os"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestNew(t *testing.T) {
	_, err := New([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConsumeHackathon(t *testing.T) {
	c, err := New([]string{"127.0.0.1:9092"}, os.Stdout)
	quit := make(chan int)
	err = c.Consume("hackathon", 0, sarama.OffsetNewest, quit)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 10)
	quit <- 1
}
