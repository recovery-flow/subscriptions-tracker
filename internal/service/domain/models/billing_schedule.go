package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BillingSchedule struct {
	UserID        uuid.UUID     `json:"user_id"`
	ScheduledDate time.Time     `json:"scheduled_date"`
	AttemptedDate *time.Time    `json:"attempted_date,omitempty"`
	Status        BillingStatus `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
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
	//case "success":
	//	return BillingStatusSuccess, nil
	case "processing":
		return BillingStatusProcessing, nil
	default:
		return "", fmt.Errorf("invalid billing status: %s", status)
	}
}
