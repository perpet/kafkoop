package kafka

import (
	"log"
	"testing"

	"github.com/Shopify/sarama"
)

func TestNewClient(t *testing.T) {

	cfg := sarama.NewConfig()
	client, err := sarama.NewClient([]string{"127.0.0.1:9092"}, cfg)
	if err != nil {
		t.Fatal(err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		t.Fatal(err)
	}

	topic := "hackathon"

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("test message"),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		t.Fatal(err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		t.Fatal(err)
	}

	pc, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		t.Fatal(err)
	}

	data := <-pc.Messages()

	log.Printf("data=%+v", *data)
}
