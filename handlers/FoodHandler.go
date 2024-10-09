package handlers

import (
	"Status418/dto"
	"Status418/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	fs services.FoodServiceInterface
}

func NewFoodHandler(fs services.FoodServiceInterface) *FoodHandler {
	return &FoodHandler{
		fs: fs,
	}
}

func (fh *FoodHandler) GetAll(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	var minimumList bool

	if err := c.ShouldBindJSON(&minimumList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foods, err := fh.fs.GetAll(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get foods"})
		return
	}

	c.JSON(http.StatusOK, foods)
}

func (fh *FoodHandler) GetByCode(c *gin.Context) {
	userId := c.Param("userId")
	// if userId == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
	// 	return
	// }
	code := c.Param("code")
	// if code == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
	// 	return
	// }

	food, err := fh.fs.GetByCode(code, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get food with code: " + code})
		return
	}
	c.JSON(http.StatusOK, food)
}

func (fh *FoodHandler) Create(c *gin.Context) {
	var newFood dto.FoodDto
	if err := c.ShouldBindJSON(&newFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	_, err := fh.fs.Create(newFood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create food item", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Food created successfully"})
}

func (fh *FoodHandler) Update(c *gin.Context) {
	var updateFood dto.FoodDto
	updateCode := c.Param("id")
	updateFood.Code = updateCode
	if err := c.ShouldBindJSON(&updateFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := fh.fs.Update(updateFood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update food item", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food updated successfully"})
}

func (fh *FoodHandler) Delete(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}

	_, err := fh.fs.Delete(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete food item", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
