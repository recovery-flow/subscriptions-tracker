package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID       `json:"id"`
	UserID          uuid.UUID       `json:"user_id"`
	PaymentMethodID uuid.UUID       `json:"payment_method_id"`
	Amount          float64         `json:"amount"`
	Currency        string          `json:"currency"`
	Status          TrnStatus       `json:"status"`
	PaymentProvider PaymentProvider `json:"payment_provider"`
	PaymentID       string          `json:"payment_id"`
	TransactionDate time.Time       `json:"transaction_date"`
}

type TrnStatus string

const (
	TrnStatusSuccess TrnStatus = "success"
	TrnStatusFailed  TrnStatus = "failed"
)

func ParseTrnStatus(s string) TrnStatus {
	switch s {
	case "success":
		return TrnStatusSuccess
	case "failed":
		return TrnStatusFailed
	default:
		return ""
	}
}

type PaymentProvider string

const (
	PaymentProviderStripe PaymentProvider = "stripe"
	PaymentProviderPaypal PaymentProvider = "paypal"
)

func ParsePaymentProvider(s string) (PaymentProvider, error) {
	switch s {
	case "stripe":
		return PaymentProviderStripe, nil
	case "paypal":
		return PaymentProviderPaypal, nil
	}
	return "", fmt.Errorf("unknown payment provider: %s", s)
}
