package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscriber struct {
	ID primitive.ObjectId `json:"id"`
}
