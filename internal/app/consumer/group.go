package consumer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func RunGroup(n int, sleep int64) {
	log.Printf("Running consumer group with %d consumers ...\n", n)
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	for i := 0; i < n; i++ {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{"localhost:9092"},
			Topic:    "wiki-test",
			GroupID:  "wiki-test-group",
			MaxBytes: 10e6, // 10MB
		})
		child, cancel := context.WithCancel(ctx)
		defer cancel()
		wg.Add(1)
		go start(child, r, &wg, time.Duration(sleep), i)
	}
	<-ctx.Done()
	wg.Wait()
	log.Printf("Done running consumer group with %d consumers\n", n)
}
