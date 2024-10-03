package dto

type PurchaseDTO struct {
	UserId    int                   
	TotalCost float64                
	Foods     []PurchaseQuantity 
}

func NewPurchaseDTO(userid int, totalcost float64, foods []PurchaseQuantity) *PurchaseDTO {
	return &PurchaseDTO{
		UserId: userid,
		TotalCost: totalcost,
		Foods: foods,
	}
}