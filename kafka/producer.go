package kafka

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func ProducerLogger(msg string, a ...interface{}) {
	args := make([]any, 0)
	for i := 0; i < len(a); i++ {
		args = append(args, fmt.Sprintf("arg%d", i), a[i])
	}
	slog.Info(msg, args...)
}

func NewProducer(brokers []string, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:        kafka.TCP(brokers...),
		Topic:       topic,
		Balancer:    &kafka.LeastBytes{},
		Logger:      kafka.LoggerFunc(ProducerLogger),
		ErrorLogger: kafka.LoggerFunc(ProducerLogger),
	}

	return &Producer{writer: writer}
}

// WriteMessage writes a single message to the topic.
func (p *Producer) WriteMessage(ctx context.Context, key, value []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

// WriteMessages writes multiple messages to the topic.
func (p *Producer) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return p.writer.WriteMessages(ctx, msgs...)
}

// Close closes the producer connection.
func (p *Producer) Close() error {
	return p.writer.Close()
}
