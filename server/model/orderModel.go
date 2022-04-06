package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Image    string             `json:"image" bson:"image"`
	Quantity int                `json:"quantity" bson:"quantity"`
	Price    int                `json:"price" bson:"price"`
	Total    int                `json:"total" bson:"total"`
}
