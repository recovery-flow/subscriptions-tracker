package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SubscriptionType struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      StatusType `json:"status"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type StatusType string

const (
	StatusTypeActive   StatusType = "active"
	StatusTypeInactive StatusType = "inactive"
)

func ParseStatusType(status string) (StatusType, error) {
	switch status {
	case "active":
		return StatusTypeActive, nil
	case "inactive":
		return StatusTypeInactive, nil
	default:
		return "", fmt.Errorf("invalid type status: %s", status)
	}
}
