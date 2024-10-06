package handlers

import "Status418/services"

type PurchaseHandler struct {
	ps services.PurchaseServiceInterface
}

func NewPurchaseHandler(ps services.PurchaseServiceInterface) *PurchaseHandler {
	return &PurchaseHandler{
		ps : ps,
	}
}

//IMPLEMENTAR LOS MÉTODOS DE LA INTERFACE PurchaseServiceInterface