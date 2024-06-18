package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Run() {
	log.Println("Running consumer...")
	// make a new reader that consumes from topic-A, partition 0
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "wiki-test",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
