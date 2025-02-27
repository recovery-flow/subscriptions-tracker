package models

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionPlan struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	BillingCycle string    `json:"billing_cycle"`
	Currency     string    `json:"currency"`
	CreatedAt    time.Time `json:"created_at"`
}
