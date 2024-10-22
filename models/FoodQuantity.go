package models

type FoodQuantity struct {
	FoodCode string `bson:"food_code"`
	Quantity int `bson:"quantity_bought"`	
}

