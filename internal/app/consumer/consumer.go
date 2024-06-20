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

// Run initiates n number of Kafka consumer that sleeps for d seconds which
// consume topic from broker.
func Run(n int, d int64, topic string, groupID string) {
	log.Printf("Running %d consumers ...\n", n)
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	for i := 0; i < n; i++ {
		var r *kafka.Reader
		if groupID == "" {
			r = newConsumer(i, topic)
		} else {
			r = newConsumerGroup(topic, groupID)
		}
		child, cancel := context.WithCancel(ctx)
		defer cancel()
		wg.Add(1)
		go start(child, r, &wg, time.Duration(d), i)
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

// newConsumer instantiates a new [kafka.Reader] that reads for partition i and
// topic of topicName and returns the pointer.
func newConsumer(i int, topicName string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topicName,
		Partition: i,
		MaxBytes:  10e6, // 10MB
	})
	return r
}

// newConsumerGroup instantiates a new [kafka.Reader] that reads for partition i and
// topic of topicName and returns the pointer. It will operate in a consumer
// group under group groupID.
func newConsumerGroup(topicName string, groupID string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topicName,
		GroupID:  groupID,
		MaxBytes: 10e6, // 10MB
	})
	return r
}
