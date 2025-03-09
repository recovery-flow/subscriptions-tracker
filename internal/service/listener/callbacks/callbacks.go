package callbacks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events"
	"github.com/segmentio/kafka-go"
)

func SubscribersPaymentEvents(ctx context.Context, m kafka.Message, evt events.InternalEvent) error {
	var ps events.SubscriptionPayment
	if err := json.Unmarshal(evt.Data, &ps); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}
	_, err := uuid.Parse(string(m.Key))
	if err != nil {
		return fmt.Errorf("failed to parse payment evemt key: %w", err)
	}
	if evt.EventType == events.SubscriptionPaymentFailedType {
		return svc.Domain.SubscriptionPaymentFailled(ctx, m.Key)
	}
	return svc.Domain.SubscriptionPaymentFailler(ctx, m.Key)
}
