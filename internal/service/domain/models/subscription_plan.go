package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SubscriptionPlan struct {
	ID                  uuid.UUID           `json:"id"`
	TypeID              uuid.UUID           `json:"type_id"`
	Price               float64             `json:"price"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	BillingInterval     int8                `json:"billing_interval"`
	BillingIntervalUnit BillingIntervalUnit `json:"billing_interval_unit"`
	Currency            string              `json:"currency"`
	Status              StatusPlan          `json:"status"`
	CreatedAt           time.Time           `json:"created_at"`
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

type BillingIntervalUnit string

const (
	Once  BillingIntervalUnit = "once"
	Day   BillingIntervalUnit = "day"
	Week  BillingIntervalUnit = "week"
	Month BillingIntervalUnit = "month"
	Year  BillingIntervalUnit = "year"
)

func ParseBillingIntervalUnit(unit string) (BillingIntervalUnit, error) {
	switch unit {
	case "once":
		return Once, nil
	case "day":
		return Day, nil
	case "week":
		return Week, nil
	case "month":
		return Month, nil
	case "year":
		return Year, nil
	default:
		return "", fmt.Errorf("invalid billing interval unit: %s", unit)
	}
}
