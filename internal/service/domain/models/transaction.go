package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID            primitive.ObjectID  `bson:"_id" json:"id"`
	UserID        *primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	PlanID        *primitive.ObjectID `bson:"plan_id,omitempty" json:"plan_id,omitempty"`
	SubID         *primitive.ObjectID `bson:"sub_id,omitempty" json:"sub_id,omitempty"`
	Amount        int                 `bson:"amount" json:"amount"`
	Currency      string              `bson:"currency" json:"currency"`
	Status        TrStatus            `bson:"status" json:"status"`
	PaymentMethod string              `bson:"payment_method" json:"payment_method"`
	ProvTxID      string              `bson:"prov_tx_id" json:"prov_tx_id"`
	CreatedAt     primitive.DateTime  `bson:"created_at" json:"created_at"`
}

type TrStatus string

const (
	TrStatusSuccess TrStatus = "success"
	TrStatusFailed  TrStatus = "failed"
)
