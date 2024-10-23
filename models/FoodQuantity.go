package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FoodQuantity struct {
	FoodCode primitive.ObjectID `bson:"_id"`
	Quantity int                `bson:"quantity_bought"`
}
