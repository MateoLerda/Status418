package dto

import "Status418/models"

type PurchaseDto struct {
	TotalCost float64
	Foods     []models.PurchaseQuantity
}


