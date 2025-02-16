package models

import (
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscriber struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	UserID uuid.UUID          `bson:"user_id" json:"user_id"`
	PlanID primitive.ObjectID `bson:"plan_id" json:"plan_id"`
	Status SubStatus          `bson:"status" json:"status"`

	StartAt   primitive.DateTime  `bson:"start_at" json:"start_at"`
	EndAt     primitive.DateTime  `bson:"end_at" json:"end_at"`
	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
}

type SubStatus string

const (
	SubStatusActive   SubStatus = "active"
	SubStatusInactive SubStatus = "inactive"
	SubStatusCanceled SubStatus = "canceled"
)

func StringToStatusSubscriber(s string) (*SubStatus, error) {
	dict := map[string]SubStatus{
		"active":   SubStatusActive,
		"inactive": SubStatusInactive,
		"canceled": SubStatusCanceled,
	}
	if _, ok := dict[s]; !ok {
		res := dict[s]
		return &res, nil
	}

	return nil, fmt.Errorf("invalid status subscriber: %s", s)
}
