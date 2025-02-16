package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionPlan struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Title    string             `bson:"title" json:"title"`
	Desc     string             `bson:"desc" json:"desc"`
	Price    int                `bson:"price" json:"price"`
	Currency string             `bson:"currency" json:"currency"`
	PayFreq  PayFreq            `bson:"pay_freq" json:"pay_freq"`
	Status   PlanStatus         `bson:"status" json:"status"`

	CanceledAt *primitive.DateTime `bson:"canceled_at,omitempty" json:"canceled_at,omitempty"`
	UpdatedAt  *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt  primitive.DateTime  `bson:"created_at" json:"created_at"`
}

type PlanStatus string

const (
	PlanStatusActive   PlanStatus = "active"
	PlanStatusInactive PlanStatus = "inactive"
	PlanStatusCanceled PlanStatus = "canceled"
)

func StringToPlanStatus(s string) (*PlanStatus, error) {
	dict := map[string]PlanStatus{
		"active":   PlanStatusActive,
		"inactive": PlanStatusInactive,
		"canceled": PlanStatusCanceled,
	}
	if _, ok := dict[s]; !ok {
		res := dict[s]
		return &res, nil
	}

	return nil, fmt.Errorf("invalid plan status: %s", s)
}

type PayFreq string

const (
	PayFreqDaily   PayFreq = "daily"
	PayFreqWeekly  PayFreq = "weekly"
	PayFreqMonthly PayFreq = "monthly"
	PayFreqYearly  PayFreq = "yearly"
)

func StringToPayFreq(s string) (*PayFreq, error) {
	dict := map[string]PayFreq{
		"daily":   PayFreqDaily,
		"weekly":  PayFreqWeekly,
		"monthly": PayFreqMonthly,
		"yearly":  PayFreqYearly,
	}
	if _, ok := dict[s]; !ok {
		res := dict[s]
		return &res, nil
	}

	return nil, fmt.Errorf("invalid pay frequency: %s", s)
}
