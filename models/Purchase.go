package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Purchase struct {
	Id_Purchase  primitive.ObjectID     `bson:"id_purchase"`
	UserId       primitive.ObjectID     `bson:"user_id"`
	PurchaseDate time.Time              `bson:"purchase_date"`
	TotalCost    float64                `bson:"total_cost"`
	Foods        []PurchaseQuantity `bson:"foods"`
}

func NewPurchase(userId primitive.ObjectID, purchaseDate time.Time, totalCost float64, foods []PurchaseQuantity) *Purchase {
	return &Purchase{
		Id_Purchase: primitive.NewObjectID(),
		UserId: userId,
		PurchaseDate: purchaseDate,
		TotalCost: totalCost,
		Foods: foods,
	}
}
