package dto

type PurchaseDto struct {
	TotalCost float64
	Foods     []PurchaseQuantity
}

func NewPurchaseDto(userId string, totalcost float64, foods []PurchaseQuantity) *PurchaseDto {
	return &PurchaseDto{

		TotalCost: totalcost,
		Foods:     foods,
	}
}
