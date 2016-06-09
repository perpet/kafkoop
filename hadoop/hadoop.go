package hadoop

import (
	"fmt"
	"os"
	"time"

	"github.com/colinmarc/hdfs"
)

type Client struct {
	c    *hdfs.Client
	path string
}

func Must(c *Client, err error) *Client {
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	return c
}

func New(dir, file string) (*Client, error) {
	if err := os.Setenv("HADOOP_USER_NAME", "hdfs"); err != nil {
		return nil, err
	}

	c, err := hdfs.New("localhost:9000")
	if err != nil {
		return nil, err
	}

	if err := c.Mkdir(dir, 755); err != nil {
		// Assume that the failure was because the directory was already
		// created.
	}

	path := dir + "/" + file

	return &Client{c, path}, nil
}

func (c *Client) Write(p []byte) (int, error) {
	fn := fmt.Sprintf("%s.%d", c.path, time.Now().UnixNano())

	w, err := c.c.Create(fn)
	if err != nil {
		return 0, err
	}

	n, err := w.Write(p)
	if err != nil {
		return 0, err
	}
	w.Close()
	return n, nil

}
