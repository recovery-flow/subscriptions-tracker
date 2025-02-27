package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	UserID          uuid.UUID  `json:"user_id"`
	PlanID          uuid.UUID  `json:"plan_id"`
	PaymentMethodID *uuid.UUID `json:"payment_method_id,omitempty"`
	Status          string     `json:"status"`
	StartDate       time.Time  `json:"start_date"`
	EndDate         time.Time  `json:"end_date"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
