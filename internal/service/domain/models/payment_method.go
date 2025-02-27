package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Type          string    `json:"type"`
	ProviderToken string    `json:"provider_token"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
}
