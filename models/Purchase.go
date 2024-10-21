package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Purchase struct {
	Id_Purchase  primitive.ObjectID `bson:"id_purchase"`
	UserCode       string 			`bson:"user_code"`
	PurchaseDate time.Time          `bson:"purchase_date"`
	TotalCost    float64            `bson:"total_cost"`
	Foods        []FoodQuantity 	`bson:"foods"`
}
