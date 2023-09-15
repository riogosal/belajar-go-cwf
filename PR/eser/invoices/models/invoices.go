package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Content struct {
	UUID  string `json:"uuid" bson:"uuid"`
	Name  string `json:"name" bson:"name"`
	Price int64  `json:"price" bson:"price"`
	Count int64  `json:"count" bson:"count"`
}

type Invoices struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID       string             `json:"uuid" bson:"uuid"`
	Code       string             `json:"code" bson:"code"`
	Alias      string             `json:"alias" bson:"alias"`
	Contents   []Content          `json:"contents" bson:"contents"`
	CreatedAt  int64              `json:"created_at" bson:"created_at"`
	TotalPrice int64              `json:"total_price" bson:"total_price"`
	TotalCount int64              `json:"total_count" bson:"total_count"`
}
