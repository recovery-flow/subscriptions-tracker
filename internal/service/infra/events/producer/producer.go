package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	SubscriptionActivate(userID, planID, typeID uuid.UUID) error
	SubscriptionDeactivate(userID, planID, typeID uuid.UUID) error

	sendMessage(topic string, event string, key string, data []byte) error
}

type producer struct {
	brokers net.Addr
	writer  *kafka.Writer
}

func NewProducer(cfg *config.Config) Producer {
	return &producer{
		brokers: kafka.TCP(cfg.Kafka.Brokers...),
		writer: &kafka.Writer{
			Addr:         kafka.TCP(cfg.Kafka.Brokers...),
			Balancer:     &kafka.LeastBytes{},
			BatchSize:    1,
			BatchTimeout: 0,
			Async:        false,
			RequiredAcks: kafka.RequireAll,
		},
	}
}

func (p *producer) SubscriptionActivate(userID, planID, typeID uuid.UUID) error {
	body, err := json.Marshal(events.SubscriptionStatus{
		PlanID:    planID.String(),
		TypeID:    typeID.String(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to marshal SubscriptionStatus event: %w", err)
	}

	return p.sendMessage(events.SubscriptionsStatusTopic, events.SubscriptionActivatedType, userID.String(), body)
}

func (p *producer) SubscriptionDeactivate(userID, planID, typeID uuid.UUID) error {
	body, err := json.Marshal(events.SubscriptionStatus{
		PlanID:    planID.String(),
		TypeID:    typeID.String(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to marshal subscription deactivate event: %w", err)
	}

	return p.sendMessage(events.SubscriptionsStatusTopic, events.SubscriptionDeactivatedType, userID.String(), body)
}

func (p *producer) sendMessage(topic string, event string, key string, body []byte) error {
	evt := events.InternalEvent{
		EventType: event,
		Data:      body,
	}
	data, err := json.Marshal(evt)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription activate event: %w", err)
	}

	msg := kafka.Message{
		Topic: topic,
		Value: data,
		Key:   []byte(key),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
