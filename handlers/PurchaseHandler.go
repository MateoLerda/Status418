package handlers

import (
	"Status418/dto"
	"Status418/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	ps services.PurchaseServiceInterface
}

func NewPurchaseHandler(ps services.PurchaseServiceInterface) *PurchaseHandler {
	return &PurchaseHandler{
		ps : ps,
	}
}

func (ph *PurchaseHandler) Create(c *gin.Context){
	userId := c.Param("userId")
	var newPurchase dto.PurchaseDto
	if err := c.ShouldBindJSON(&newPurchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	purchase, err := ph.ps.Create(userId, newPurchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase", "details": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, purchase)

} 

