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
		rs: rs,
	}
}

func (rh *RecipeHandler) GetAll(c *gin.Context) {
	userId := c.Param("userId")
	var filters dto.FiltersDto
	err := c.ShouldBindJSON(&filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get filters"})
	}

	recipes, err := rh.rs.GetAll(userId, filters)

	if err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recipes from database",
		})
		return
	}

	if err.Error() == "nocontent" {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Not found any recipe",
		})
		return
	}
	c.JSON(http.StatusOK, recipes)
}
func (rh *RecipeHandler) Create(c *gin.Context) {
	var recipe dto.RecipeDto
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := rh.rs.Create(recipe)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe" + err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func (rh *RecipeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	res, err := rh.rs.Delete(id)

	if err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe" + id})
		return
	}

	if err.Error() == "notfound" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found any recipe with id:" + id})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (rh *RecipeHandler) Update(c *gin.Context) {
	var recipe dto.RecipeDto

	err:= c.ShouldBindJSON(&recipe)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := rh.rs.Update(recipe)

	if err.Error() == "internal"{
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete recipe"})
	}
	if err.Error() == "notfound"{
		c.JSON(http.StatusNotFound,  gin.H{"error": "Not found any recipe"})
	}

	c.JSON(http.StatusOK, res)
}
