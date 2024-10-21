package handlers

import (
	"Status418/dto"
	"Status418/services"
	"Status418/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	foodService services.FoodServiceInterface
}

func NewFoodHandler(foodService services.FoodServiceInterface) *FoodHandler {
	return &FoodHandler{
		foodService: foodService,
	}
}

func (foodHandler *FoodHandler) GetAll(c *gin.Context) {
	userCode := utils.GetUserInfoFromContext(c).Code
	if userCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
		return
	}
	minimumList, _ := strconv.ParseBool(c.Query("minimumList"))
	
	foods, err := foodHandler.foodService.GetAll(userCode, minimumList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foods)
}

func (foodHandler *FoodHandler) GetByCode(c *gin.Context) {
	userCode := utils.GetUserInfoFromContext(c).Code
	if userCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
		return
	}
	foodCode:= c.Param("foodCode")
	food, err := foodHandler.foodService.GetByCode(foodCode, userCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, food)
}

func (foodHandler *FoodHandler) Create(c *gin.Context) {
	var newFood dto.FoodDto
	if err := c.ShouldBindJSON(&newFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	insertedId, err := foodHandler.foodService.Create(newFood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create food item", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Food created successfully", "details": insertedId})
}

func (foodHandler *FoodHandler) Update(c *gin.Context) {
	var updateFood dto.FoodDto
	updateCode := c.Param("foodCode")
	if err := c.ShouldBindJSON(&updateFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateFood.Code = updateCode

	_, err := foodHandler.foodService.Update(updateFood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food item", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food updated successfully"})
}

func (foodHandler *FoodHandler) Delete(c *gin.Context) {
	deleteCode := c.Param("foodCode")

	_, err := foodHandler.foodService.Delete(deleteCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete food item", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
