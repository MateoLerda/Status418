package dto

type PurchaseDto struct {
	TotalCost float64           `json:"total_cost" validate:"gte=0.0"`
	Foods     []FoodQuantityDTO `json:"foods"`
}
