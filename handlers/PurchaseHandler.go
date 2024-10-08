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
		c.JSON(422, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	purchase, err := ph.ps.Create(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create purchase", "details": err.Error()})
		return
	}
	c.JSON(201, purchase)

} 

func (ph *PurchaseHandler) GetFoodWithQuantityLessThanMinimum(c *gin.Context){
	userId := c.Param("userId")
	food, err := ph.ps.GetFoodWithQuantityLessThanMinimum(userId)
	if err.Error()=="internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "An internal server error has ocurred"})
		return
	}
	
	if err.Error() == "nocontent"{
		c.JSON(http.StatusNoContent, gin.H{"error": "Not found any foods with quantity less than minimum"})
		return
	}
	c.JSON(200, food)

}