package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PaymentMethod struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Type          PayType   `json:"type"`
	ProviderToken string    `json:"provider_token"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PayType string

const (
	PayTypeGoogle  PayType = "paypal"
	PayTypeApple   PayType = "apple_pay"
	PayTypeSamsung PayType = "google_pay"
	PayTypePaypal  PayType = "samsung_pay"
)

func ParsePayType(s string) (PayType, error) {
	switch s {
	case "google":
		return PayTypeGoogle, nil
	case "apple":
		return PayTypeApple, nil
	case "samsung":
		return PayTypeSamsung, nil
	case "paypal":
		return PayTypePaypal, nil
	default:
		return "", fmt.Errorf("unknown pay type: %s", s)
	}
}
