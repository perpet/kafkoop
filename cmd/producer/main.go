package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/perpet/kafkoop/kafka/producer"
)

var inputFile string

func init() {
	flag.StringVar(&inputFile, "f", "", "input file")
}

func main() {
	flag.Parse()

	if inputFile == "" {
		fmt.Println("-f argument required")
		os.Exit(1)
	}

	r, err := producer.NewLineReader(inputFile)
	if err != nil {
		fmt.Printf("error: could create Reader from '%s': %s",
			inputFile, err)
		os.Exit(1)
	}

	producer, err := producer.New([]string{"127.0.0.1:9092"}, r)
	if err != nil {
		fmt.Printf("error: create Producer: '%s'", err)
		os.Exit(1)

	}

	producer.Produce()

}
