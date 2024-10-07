package dto

import "Status418/models"

type PurchaseDto struct {
	TotalCost float64
	Foods     []models.PurchaseQuantity
}

func NewPurchaseDto(userId string, totalCost float64, foods []models.PurchaseQuantity) *PurchaseDto {
	return &PurchaseDto{

		TotalCost: totalCost,
		Foods:     foods,
	}
}




