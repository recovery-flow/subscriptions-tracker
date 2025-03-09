package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BillingSchedule struct {
	UserID        uuid.UUID             `json:"user_id"`
	SchedulesDate time.Time             `json:"schedules_date"`
	AttemptedDate *time.Time            `json:"attempted_date,omitempty"`
	Status        ScheduleBillingStatus `json:"status"`
	UpdatedAt     time.Time             `json:"updated_at"`
	CreatedAt     time.Time             `json:"created_at"`
}

type ScheduleBillingStatus string

const (
	ScheduleBillingStatusPlanned    ScheduleBillingStatus = "planned"
	ScheduleBillingStatusFailed     ScheduleBillingStatus = "failed"
	ScheduleBillingStatusProcessing ScheduleBillingStatus = "processing"
)

func ParseBillingStatus(status string) (ScheduleBillingStatus, error) {
	switch status {
	case "planned":
		return ScheduleBillingStatusPlanned, nil
	case "failed":
		return ScheduleBillingStatusFailed, nil
	case "processing":
		return ScheduleBillingStatusProcessing, nil
	default:
		return "", fmt.Errorf("invalid billing status: %s", status)
	}
}
