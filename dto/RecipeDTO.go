package dto

import (
	"Status418/enums"
	"Status418/models"
	"Status418/utils"
)

type RecipeDto struct {
	Id          string
	Name        string
	Ingredients []FoodDto
	Moment      enums.Moment
	Description string
	UserCode      string
}

func NewRecipeDto(model models.Recipe) *RecipeDto {
	var dtoIngredients []FoodDto

	for _, food := range model.Ingredients {
		dtoIngredients = append(dtoIngredients, *NewFoodDto(food))
	}

	return &RecipeDto{
		Id: utils.GetStringIDFromObjectID(model.Id),
		Name:        model.Name,
		Ingredients: dtoIngredients,
		Moment: model.Moment,
		Description: model.Description,
		UserCode:      model.UserCode,
	}
}

func (dto RecipeDto) GetModel() models.Recipe {
	var ingredients []models.Food

	for _, food := range dto.Ingredients {
		ingredients = append(ingredients, food.GetModel())
	}

	return models.Recipe{
		Id: utils.GetObjectIDFromStringID(dto.Id),
		Name: dto.Name,
		Ingredients: ingredients,
		Moment: dto.Moment,
		Description: dto.Description,
		UserCode: dto.UserCode,
	}
}
