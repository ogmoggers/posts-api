package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Producer struct {
	writer *kafka.Writer
	topic  string
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{
		writer: writer,
		topic:  topic,
	}, nil
}

func (p *Producer) PublishPostCreated(post Post) error {
	value, err := json.Marshal(post)
	if err != nil {
		return err
	}

	headers := []kafka.Header{
		{Key: "event", Value: []byte("post_created")},
	}

	return p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value:   value,
			Headers: headers,
		},
	)
}

func (p *Producer) PublishPostUpdated(post Post) error {
	value, err := json.Marshal(post)
	if err != nil {
		return err
	}

	headers := []kafka.Header{
		{Key: "event", Value: []byte("post_updated")},
	}

	return p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value:   value,
			Headers: headers,
		},
	)
}

func (p *Producer) PublishPostDeleted(postID string) error {
	headers := []kafka.Header{
		{Key: "event", Value: []byte("post_deleted")},
	}

	return p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value:   []byte(postID),
			Headers: headers,
		},
	)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}