package producer

import "testing"

func TestNew(t *testing.T) {
	_, err := New([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		t.Fatal(err)
	}
}
