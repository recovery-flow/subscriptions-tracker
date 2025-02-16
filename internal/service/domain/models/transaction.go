package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID            primitive.ObjectID  `bson:"_id" json:"id"`
	UserID        *primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
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

func StringToTrStatus(s string) (*TrStatus, error) {
	dict := map[string]TrStatus{
		"success": TrStatusSuccess,
		"failed":  TrStatusFailed,
	}
	if _, ok := dict[s]; !ok {
		res := dict[s]
		return &res, nil
	}

	return nil, fmt.Errorf("invalid transaction status: %s", s)
}
