package events

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KafkaBus struct {
	writer *kafka.Writer
}

func NewKafkaBus(brokers, topic string) *KafkaBus {
	return &KafkaBus{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(brokers),
			Topic: topic,
		},
	}
}

func (b *KafkaBus) Publish(event any) {
	data, _ := json.Marshal(event)

	_ = b.writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: data,
		},
	)
}
