package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "my-group",
		Topic:    "my-topic",
		MaxBytes: 10, // 10 MB
	}

	reader := kafka.NewReader(conf)
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Printf("message at offset %d: %s = %s\n", msg.Offset, msg.Key, msg.Value)
	} // loop forever

}
