package evebody

import "time"

type CreateSubscription struct {
	UserID         string    `json:"user_id"`
	SubscriptionID string    `json:"subscription_id"`
	PlanID         string    `json:"plan_id"`
	TypeID         string    `json:"type_id"`
	CreatedAt      time.Time `json:"created_at"`
}
