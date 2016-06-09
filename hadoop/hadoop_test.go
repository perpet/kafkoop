package hadoop

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New("/tmp", "foo")
	if err != nil {
		t.Fatalf("New failed: %s", err)
	}
}

func TestWrite(t *testing.T) {
	c, err := New("/tmp", "foo")
	if err != nil {
		t.Fatalf("New failed: %s", err)
	}

	str := fmt.Sprintf("%s\n", "dummy string")

	n, err := c.Write([]byte(str))
	if err != nil {
		t.Errorf("Write  failed: %s", err)
	}
	if n != len(str) {
		t.Fatalf("expected n=%d, got n=%d", len(str), n)
	}
}
