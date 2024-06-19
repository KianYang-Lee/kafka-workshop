package consumer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func Run(n int, sleep int64) {
	log.Printf("Running %d consumers ...\n", n)
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	for i := 0; i < n; i++ {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{"localhost:9092"},
			Topic:     "wiki-test",
			Partition: i,
			MaxBytes:  10e6, // 10MB
		})
		child, cancel := context.WithCancel(ctx)
		defer cancel()
		wg.Add(1)
		go start(child, r, &wg, time.Duration(sleep), i)
	}
	<-ctx.Done()
	wg.Wait()
	log.Printf("Done running %d consumers\n", n)
}

// start runs a blocking operation as a Kafka consumer.
func start(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, d time.Duration, id int) {
	log.Printf("Running consumer ID %d with sleep of %d sec...\n", id, d)
	defer wg.Done()

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
			m, err := r.ReadMessage(ctx)
			if err != nil {
				break
			}
			fmt.Printf("Consumer ID %d: message at offset %d: %s = %s\n", id, m.Offset, string(m.Key), string(m.Value))
			if d > 0 {
				time.Sleep(time.Duration(d) * time.Second)
			}
		}
	}
	fmt.Println("close")
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
