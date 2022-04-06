package model

type Item struct {
	Name        string `json:"name" bson:"name"`
	Image       string `json:"image" bson:"image"`
	Description string `json:"des" bson:"des"`
	Quantity    int    `json:"quantity" bson:"quantity"`
	Price       int    `json:"price" bson:"price"`
}
