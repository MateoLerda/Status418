package handlers

import (
	"Status418/dto"
	"Status418/services"
	"Status418/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	recipeService services.RecipeServiceInterface
}

func NewRecipeHandler(recipeService services.RecipeServiceInterface) *RecipeHandler {
	return &RecipeHandler{
		recipeService: recipeService,
	}
}

func (recipeHandler *RecipeHandler) GetAll(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	if user.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
		return
	}
	var filters dto.FiltersDto
	filters.Aproximation = c.Query("filter_aproximation")
	filters.Moment = c.Query("filter_moment")
	filters.Type = c.Query("filter_type")
	filters.All, _ = strconv.ParseBool(c.Query("filter_all"))

	recipes, err := recipeHandler.recipeService.GetAll(user.Code, filters)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recipes from database",
		})
		return
	}

	if err != nil && err.Error() == "nocontent" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Not found any recipe",
		})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func (recipeHandler *RecipeHandler) Create(c *gin.Context) {
	var recipe dto.RecipeDto
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := utils.GetUserInfoFromContext(c)
	recipe.UserCode = user.Code
	if recipe.UserCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
		return
	}
	res, err := recipeHandler.recipeService.Create(recipe)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create recipe " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (recipeHandler *RecipeHandler) Delete(c *gin.Context) {
	id := c.Param("recipeid")
	_, err := recipeHandler.recipeService.Delete(id)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe with id: " + id})
		return
	}

	if err != nil && err.Error() == "notfound" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found any recipe with id: " + id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted recipe with id: " + id})

}

func (recipeHandler *RecipeHandler) Update(c *gin.Context) {
	id := c.Param("recipeid")
	var recipe dto.RecipeDto
	err := c.ShouldBindJSON(&recipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe.Id = id
	res, err := recipeHandler.recipeService.Update(recipe)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to delete recipe"})
	}
	if err != nil && err.Error() == "notfound" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found any recipe with id: " + id})
	}

	c.JSON(http.StatusOK, res)
}

func (recipeHandler *RecipeHandler) Cook(c *gin.Context) {
	recipeId := c.Param("recipe_id") //ver que onda
	recipeObjectId := utils.GetObjectIDFromStringID(recipeId)
	userInfo := utils.GetUserInfoFromContext(c)

	if userInfo.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userInfo is required"})
		return
	}
	cancel, _ := strconv.ParseBool(c.Query("cancel"))
	res, err := recipeHandler.recipeService.Cook(userInfo.Code, recipeObjectId, cancel)

	if err != nil && err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel recipe"})
	}
	if err != nil && err.Error() == "notfound" {
		c.JSON(http.StatusOK, gin.H{"message": "Not found any recipe with id: " + recipeId})
	}

	c.JSON(http.StatusOK, res)
}
