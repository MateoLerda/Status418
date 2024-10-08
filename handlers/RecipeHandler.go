package handlers

import (
	"Status418/dto"
	"Status418/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	rs services.RecipeServiceInterface
}

func NewRecipeHandler(rs services.RecipeServiceInterface) *RecipeHandler {
	return &RecipeHandler{
		rs : rs,
	}
}

func (rh *RecipeHandler) GetAll(c *gin.Context){
	userId := c.Param("userId")
	var filters dto.FiltersDto
	err := c.ShouldBindJSON(&filters)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get filters", "Moment": filters.Moment})
	}

	recipes, err := rh.rs.GetAll(userId,filters)

	if err.Error()== "internal"{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recipes from database",
		})
		return
	}

	if err.Error() == "nocontent"{
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Not found any recipe",
		})
		return
	}
	c.JSON(http.StatusOK, recipes)
}
func (rh *RecipeHandler) Create(c *gin.Context){

} 

func (rh *RecipeHandler) Delete(c *gin.Context){

}

func (rh *RecipeHandler) Update(c *gin.Context){

}
