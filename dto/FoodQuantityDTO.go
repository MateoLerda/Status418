package dto

type FoodQuantityDTO struct {
	FoodCode string `json:"_id"`
	Quantity int    `json:"quantity_bought" validate:"numeric"`
}
