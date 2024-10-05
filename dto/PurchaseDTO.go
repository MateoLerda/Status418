package dto

type PurchaseDto struct {
	TotalCost float64
	Foods     []PurchaseQuantity
}

func NewPurchaseDto(userId string, totalCost float64, foods []PurchaseQuantity) *PurchaseDto {
	return &PurchaseDto{

		TotalCost: totalCost,
		Foods:     foods,
	}
}
