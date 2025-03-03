package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	UserID          uuid.UUID         `json:"user_id"`
	PlanID          uuid.UUID         `json:"plan_id"`
	PaymentMethodID uuid.UUID         `json:"payment_method_id"`
	State           SubscriptionState `json:"state"`
	Availability    PlanAvailability  `json:"availability"`
	StartDate       time.Time         `json:"start_date"`
	EndDate         time.Time         `json:"end_date"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

type SubscriptionState string

const (
	SubscriptionStatusActive   SubscriptionState = "active"
	SubscriptionStatusInactive SubscriptionState = "inactive"
	SubscriptionStatusExpired  SubscriptionState = "expired"
)

func ParseSubscriptionState(status string) (SubscriptionState, error) {
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

type PlanAvailability string

const (
	PlanAvailable  PlanAvailability = "available"  // Подписка доступна для покупки
	PlanDeprecated PlanAvailability = "deprecated" // Больше не доступна для новых пользователей
	PlanRemoved    PlanAvailability = "removed"    // Полностью удалена из системы
)

func ParsePlanAvailability(availability string) (PlanAvailability, error) {
	switch availability {
	case "available":
		return PlanAvailable, nil
	case "deprecated":
		return PlanDeprecated, nil
	case "removed":
		return PlanRemoved, nil
	default:
		return "", fmt.Errorf("invalid plan availability: %s", availability)
	}
}
