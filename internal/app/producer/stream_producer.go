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

func RunStream(streamName string) {
	fmt.Println("Running stream producer...")
	wg := sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	stream := make(chan *sse.Event)
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    streamName,
		Balancer: &kafka.LeastBytes{},
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
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
			fmt.Printf("Count %d: Written msg: [%s]\n\n", count, string(msg.Data))
			count++
		}
	}()

	client := sse.NewClient(fmt.Sprintf("https://stream.wikimedia.org/v2/stream/%s", streamName))

	wg.Add(1)
	go func() {
		client.SubscribeChanRawWithContext(ctx, stream)
		wg.Done()
	}()
	<-ctx.Done()
	client.Unsubscribe(stream)
	close(stream)
	wg.Wait()
	fmt.Println("Done running stream producer")
}
