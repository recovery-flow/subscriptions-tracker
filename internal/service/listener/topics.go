package listener

import (
	"context"

	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/listener/callbacks"
	"github.com/segmentio/kafka-go"
)

type TopicConfig struct {
	Topic      string
	ReplyTopic string
	Callback   func(ctx context.Context, m kafka.Message, evt events.InternalEvent) error
	OnSuccess  func(ctx context.Context, m kafka.Message, ie events.InternalEvent) error
	OnError    func(ctx context.Context, m kafka.Message, ie events.InternalEvent, err error)
}

var TopicsConfig = []TopicConfig{
	{
		Topic:    events.SubscriptionPaymentsTopic,
		Callback: callbacks.SubscribersPaymentEvents,
	},
}
