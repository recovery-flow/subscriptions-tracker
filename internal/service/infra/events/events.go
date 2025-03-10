package events

import (
	"encoding/json"
	"time"
)

const (
	SubscriptionsStatusTopic    = "subscriptions"
	SubscriptionActivatedType   = "subscription_activated"
	SubscriptionDeactivatedType = "subscription_deactivated"

	SubscriptionPaymentsTopic      = "subscription_payments"
	SubscriptionPaymentSuccessType = "subscription_payment_success"
	SubscriptionPaymentFailedType  = "subscription_payment_failed"
)

type InternalEvent struct {
	EventType string          `json:"event_type"`
	Data      json.RawMessage `json:"data"`
}

type SubscriptionStatus struct {
	PlanID    string    `json:"plan_id"`
	TypeID    string    `json:"type_id"`
	CreatedAt time.Time `json:"created_at"`
}

type SubscriptionPayment struct {
	PlanID    string    `json:"plan_id"`
	TypeID    string    `json:"type_id"`
	Amount    *float64  `json:"amount,omitempty"`
	Currency  *string   `json:"currency,omitempty"`
	Reason    *string   `json:"reason,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
