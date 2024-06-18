package consumer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/segmentio/kafka-go"
)

func RunGroups(n int, sleep int64) {
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
		child = context.WithValue(child, "sleepDuration", sleep)
		child = context.WithValue(child, "consumerID", i)
		go start(child, r, &wg)
		wg.Add(1)
	}
	<-ctx.Done()
	wg.Wait()
	log.Printf("Done running consumer group with %d consumers\n", n)
}
