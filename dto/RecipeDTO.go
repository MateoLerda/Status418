package dto

import (
	"Status418/enums"
	"Status418/models"
	"Status418/utils"
)

type RecipeDto struct {
	Id          string
	Name        string
	Ingredients []FoodQuantityDTO
	Moment      enums.Moment
	Description string
	UserCode    string
}

func NewRecipeDto(model models.Recipe) *RecipeDto {
	var dtoIngredients []FoodQuantityDTO

	for _, food := range model.Ingredients {
		dtoIngredients = append(dtoIngredients, FoodQuantityDTO{FoodCode: food.FoodCode, Quantity: food.Quantity})
	}

	return &RecipeDto{
		Id:          utils.GetStringIDFromObjectID(model.Id),
		Name:        model.Name,
		Ingredients: dtoIngredients,
		Moment:      model.Moment,
		Description: model.Description,
		UserCode:    model.UserCode,
	}
}


func (dto RecipeDto) GetModel() models.Recipe {
	var ingredients []models.FoodQuantity

	for _, food := range dto.Ingredients {
		ingredients = append(ingredients, models.FoodQuantity {
			FoodCode: food.FoodCode,		
			Quantity: food.Quantity,
		})
	}

	return models.Recipe{
		Id:          utils.GetObjectIDFromStringID(dto.Id),
		Name:        dto.Name,
		Ingredients: ingredients,
		Moment:      dto.Moment,
		Description: dto.Description,
		UserCode:    dto.UserCode,
	}
}
