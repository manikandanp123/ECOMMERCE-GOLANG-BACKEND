package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Phone    int                `json:"phone" bson:"phone"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}
