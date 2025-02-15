package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscriber struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       uuid.UUID          `bson:"user_id" json:"user_id"`
	PlanID       primitive.ObjectID `bson:"plan_id" json:"plan_id"`
	StreakMonths int                `bson:"streak_months" json:"streak_months"`
	Status       string             `bson:"status" json:"status"`

	StartAt   primitive.DateTime `bson:"start_at" json:"start_at"`
	EndAt     primitive.DateTime `bson:"end_at" json:"end_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
}
