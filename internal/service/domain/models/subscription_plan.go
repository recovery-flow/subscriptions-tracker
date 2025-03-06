package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SubscriptionPlan struct {
	ID              uuid.UUID    `json:"id"`
	TypeID          uuid.UUID    `json:"type_id"`
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Price           float64      `json:"price"`
	Currency        string       `json:"currency"`
	BillingInterval int8         `json:"billing_interval"`
	BillingCycle    BillingCycle `json:"billing_interval_unit"`
	Status          StatusPlan   `json:"status"`
	UpdatedAt       time.Time    `json:"updated_at"`
	CreatedAt       time.Time    `json:"created_at"`
}

type StatusPlan string

const (
	StatusPlanActive   StatusPlan = "active"
	StatusPlanInactive StatusPlan = "inactive"
)

func ParseStatusPlan(status string) (StatusPlan, error) {
	switch status {
	case "active":
		return StatusPlanActive, nil
	case "inactive":
		return StatusPlanInactive, nil
	default:
		return "", fmt.Errorf("invalid plan status: %s", status)
	}
}

type BillingCycle string

const (
	CycleOnce  BillingCycle = "once"
	CycleDay   BillingCycle = "day"
	CycleWeek  BillingCycle = "week"
	CycleMonth BillingCycle = "month"
	CycleYear  BillingCycle = "year"
)

func ParseBillingCycle(unit string) (BillingCycle, error) {
	switch unit {
	case "once":
		return CycleOnce, nil
	case "day":
		return CycleDay, nil
	case "week":
		return CycleWeek, nil
	case "month":
		return CycleMonth, nil
	case "year":
		return CycleYear, nil
	default:
		return "", fmt.Errorf("invalid billing interval unit: %s", unit)
	}
}
