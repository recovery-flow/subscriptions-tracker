package models

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionCancellation struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"user_id"`
	CancellationDate time.Time `json:"cancellation_date"`
	Reason           *string   `json:"reason,omitempty"`
}
