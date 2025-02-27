package models

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionTransaction struct {
	ID              uuid.UUID  `json:"id"`
	UserID          uuid.UUID  `json:"user_id"`
	PaymentMethodID *uuid.UUID `json:"payment_method_id,omitempty"`
	Amount          float64    `json:"amount"`
	Currency        string     `json:"currency"`
	Status          string     `json:"status"`
	PaymentProvider string     `json:"payment_provider"`
	PaymentID       *string    `json:"payment_id,omitempty"`
	TransactionDate time.Time  `json:"transaction_date"`
}
