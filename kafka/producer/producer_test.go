package producer

import (
	"os"
	"testing"
)

var inputFile string

func TestMain(m *testing.M) {
	inputFile = os.Getenv("KAFKA_INPUT")
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	_, err := New([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLineReader(t *testing.T) {
	_, err := NewLineReader(inputFile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProducer(t *testing.T) {
	r, _ := NewLineReader(inputFile)

	producer, err := New([]string{"127.0.0.1:9092"}, r)
	if err != nil {
		t.Fatal(err)
	}

	err = producer.Produce()
	if err != nil {
		t.Fatal(err)
	}
}
