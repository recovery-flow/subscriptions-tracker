package callbacks

import (
	"context"

	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events"
	"github.com/segmentio/kafka-go"
)

func SubscribersPaymentEvents(ctx context.Context, svc *service.Service, m kafka.Message, evt events.InternalEvent) error {
	//var ps events.SubscriptionPayment
	//if err := json.Unmarshal(evt.Data, &ps); err != nil {
	//	return fmt.Errorf("failed to unmarshal event: %w", err)
	//}
	//_, err := uuid.Parse(string(m.Key))
	//if err != nil {
	//	return fmt.Errorf("failed to parse payment evemt key: %w", err)
	//}
	//if evt.EventType == events.SubscriptionPaymentFailedType {
	//	return domain.SubscriptionPaymentFailled(ctx, m.Key)
	//}
	//return domain.SubscriptionPaymentFailler(ctx, m.Key)
	return nil
}
