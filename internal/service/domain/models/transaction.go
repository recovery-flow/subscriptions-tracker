package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	PlanID        primitive.ObjectID `bson:"plan_id" json:"plan_id"`
	SubID         primitive.ObjectID `bson:"sub_id" json:"sub_id"`
	Amount        int                `bson:"amount" json:"amount"`
	Currency      string             `bson:"currency" json:"currency"`
	PaymentMethod string             `bson:"payment_method" json:"payment_method"`
	ProvTxID      string             `bson:"prov_tx_id" json:"prov_tx_id"`
	CreatedAt     primitive.DateTime `bson:"created_at" json:"created_at"`
}
