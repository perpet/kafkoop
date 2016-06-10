package main

import (
	"log"
	"os"
	"strings"

	"github.com/perpet/kafkoop/kafka/consumer"

	"github.com/Shopify/sarama"
	ui "github.com/gizak/termui"
)

type consoleBox struct {
	par  *ui.Par
	buf  []string
	rows int
}

func (c *consoleBox) Write(p []byte) (n int, err error) {
	c.buf = append(c.buf[1:], string(p))
	c.par.Text = strings.Join(c.buf, "\n")

	ui.Render(ui.Body)

	return len(p), nil
}

func newConsoleBox(title string, rows int) *consoleBox {
	c := new(consoleBox)
	c.rows = rows
	c.buf = make([]string, c.rows)

	c.par = ui.NewPar("")
	c.par.BorderLabel = " " + title + " "
	c.par.BorderFg = ui.ColorYellow
	c.par.Height = 2 + c.rows

	return c
}

func main() {
	err := ui.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	for _, topic := range os.Args[1:] {
		box := newConsoleBox("Topic: "+topic, 3)

		c, err := consumer.New([]string{"127.0.0.1:9092"}, box)
		if err != nil {
			log.Fatal(err)
		}

		quit := make(chan int)

		c.Consume(topic, 0, sarama.OffsetNewest, quit)

		ui.Body.AddRows(
			ui.NewRow(
				ui.NewCol(12, 0, box.par)))
	}

	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop()
}
