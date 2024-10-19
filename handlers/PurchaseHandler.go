package handlers

import (
	"Status418/dto"
	"Status418/services"
	"Status418/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	purchaseService services.PurchaseServiceInterface
}

func NewPurchaseHandler(purchaseService services.PurchaseServiceInterface) *PurchaseHandler {
	return &PurchaseHandler{
		purchaseService: purchaseService,
	}
}

func (purchaseHandler *PurchaseHandler) Create(c *gin.Context){
	userCode := (dto.NewUser(utils.GetUserInfoFromContext(c))).Code
		
	var newPurchase dto.PurchaseDto
	if err := c.ShouldBindJSON(&newPurchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	purchase, err := purchaseHandler.purchaseService.Create(userCode, newPurchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase", "details": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, purchase)

} 

