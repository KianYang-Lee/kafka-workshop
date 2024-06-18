package producer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/r3labs/sse/v2"
	"github.com/segmentio/kafka-go"
)

func RunStream() {
	fmt.Println("Running stream producer...")
	wg := sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	stream := make(chan *sse.Event)
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "wiki-test",
		Balancer: &kafka.LeastBytes{},
	}

	go func() {
		for msg := range stream {
			err := w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("test"),
					Value: []byte(msg.Data),
				},
			)
			if err != nil {
				fmt.Println(err)
				stop()
				break
			}
		}
		wg.Done()
	}()
	wg.Add(1)

	client := sse.NewClient("https://stream.wikimedia.org/v2/stream/test")

	go client.SubscribeChanRawWithContext(ctx, stream)
	wg.Add(1)

	<-ctx.Done()
	client.Unsubscribe(stream)
	close(stream)
	wg.Done()
	wg.Wait()
	fmt.Println("Done running stream producer")
}
