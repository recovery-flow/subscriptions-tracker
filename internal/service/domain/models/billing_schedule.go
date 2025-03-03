package models

import (
	"time"

	"github.com/google/uuid"
)

type BillingSchedule struct {
	ID            uuid.UUID     `json:"id"`
	UserID        uuid.UUID     `json:"user_id"`
	ScheduledDate time.Time     `json:"scheduled_date"`
	AttemptedDate *time.Time    `json:"attempted_date,omitempty"`
	Status        BillingStatus `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
}

type BillingStatus string

const (
	BillingStatusPlanned BillingStatus = "planned"
	BillingStatusFailed  BillingStatus = "failed"
	BillingStatusPaid    BillingStatus = "success"
)
