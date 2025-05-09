package producer

import (
	"context"
	"encoding/json"
	"fmt"
	Nats "github.com/bwjson/kolesa_api/pkg/nats"
	"log"
	"time"
)

const (
	PushTimeout       = 30 * time.Second
	CarCreatedSubject = "car.created"
)

type CarProducer struct {
	natsClient *Nats.NatsClient
	subject    string
}

func NewCarProducer(
	natsClient *Nats.NatsClient,
) *CarProducer {
	return &CarProducer{
		natsClient: natsClient,
		subject:    CarCreatedSubject,
	}
}

type CarCreatedEvent struct {
	CarID  string `json:"car_id"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Brand  string `json:"brand"`
	Model  string `json:"model"`
}

func (c *CarProducer) Push(ctx context.Context, event CarCreatedEvent) error {
	log.Println("NATS: START")
	ctx, cancel := context.WithTimeout(ctx, PushTimeout)
	defer cancel()

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = c.natsClient.Conn.Publish(c.subject, data)
	if err != nil {
		return fmt.Errorf("nats publish: %w", err)
	}

	log.Println("Published car.created event to NATS")
	return nil
}
