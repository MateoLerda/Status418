package dto

type PurchaseQuantity struct {
	FoodCode int `bson:"food_code"`
	Quantity int `bson:"quantity_bought"`	
}

