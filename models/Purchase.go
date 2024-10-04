package models

import (
	"Status418/dto"
	"time"
)

type Purchase struct {
	Id_Purchase  int                    `bson:"id_purchase"`
	UserId       int                    `bson:"user_id"`
	PurchaseDate time.Time              `bson:"purchase_date"`
	TotalCost    float64                `bson:"total_cost"`
	Foods        []dto.PurchaseQuantity `bson:"foods"`
}
