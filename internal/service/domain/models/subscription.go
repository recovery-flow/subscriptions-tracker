package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	UserID          uuid.UUID          `json:"user_id"`
	PlanID          uuid.UUID          `json:"plan_id"`
	PaymentMethodID uuid.UUID          `json:"payment_method_id"`
	Status          SubscriptionStatus `json:"status"`
	StartDate       time.Time          `json:"start_date"`
	EndDate         time.Time          `json:"end_date"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusInactive SubscriptionStatus = "inactive"
	SubscriptionStatusExpired  SubscriptionStatus = "expired"
)

func ParseSubscriptionStatus(status string) (SubscriptionStatus, error) {
	switch status {
	case "active":
		return SubscriptionStatusActive, nil
	case "inactive":
		return SubscriptionStatusInactive, nil
	case "expired":
		return SubscriptionStatusExpired, nil
	default:
		return "", fmt.Errorf("invalid subscription status: %s", status)
	}
}
