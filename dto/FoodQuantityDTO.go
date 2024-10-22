package dto

type FoodQuantityDTO struct {
	FoodCode string `json:"food_code"`
	Quantity int    `json:"quantity_bought" validate:"numeric"`
}
