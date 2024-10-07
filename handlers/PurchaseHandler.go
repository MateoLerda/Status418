package handlers

import (
	"Status418/dto"
	"Status418/services"
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	purchase, err := ph.ps.Create(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, purchase)

} 

func (ph *PurchaseHandler) GetFoodWithQuantityLessThanMinimum(c *gin.Context){
	userId := c.Param("userId")
	food, err := ph.ps.GetFoodWithQuantityLessThanMinimum(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if len(*food) == 0 {
		c.JSON(200, gin.H{"message": "No food items with quantity less than minimum"})
		return
	}
	c.JSON(200, food)

}