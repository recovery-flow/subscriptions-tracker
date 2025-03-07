package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BillingSchedule struct {
	UserID        uuid.UUID     `json:"user_id"`
	SchedulesDate time.Time     `json:"schedules_date"`
	AttemptedDate *time.Time    `json:"attempted_date,omitempty"`
	Status        BillingStatus `json:"status"`
	UpdatedAt     time.Time     `json:"updated_at"`
	CreatedAt     time.Time     `json:"created_at"`
}

type BillingStatus string

const (
	BillingStatusPlanned    BillingStatus = "planned"
	BillingStatusFailed     BillingStatus = "failed"
	BillingStatusProcessing BillingStatus = "processing"
)

func ParseBillingStatus(status string) (BillingStatus, error) {
	switch status {
	case "planned":
		return BillingStatusPlanned, nil
	case "failed":
		return BillingStatusFailed, nil
	case "processing":
		return BillingStatusProcessing, nil
	default:
		return "", fmt.Errorf("invalid billing status: %s", status)
	}
}
