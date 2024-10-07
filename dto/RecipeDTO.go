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
	UserId      string
}

func NewRecipeDto(model models.Recipe) *RecipeDto {
	var dtoIngredients []FoodDto

	for i, food := range model.Ingredients {
		dtoIngredients = append(dtoIngredients, NewFoodDto(model.Ingredients[i]))
	}

	return &RecipeDto{
		Id: utils.GetStringIDFromObjectID(model.Id),
		Name:        model.Name,
		Ingredients: dtoIngredients,
		Moment: model.Moment,
		Description: model.Description,
		UserId:      utils.GetStringIDFromObjectID(model.UserId),
	}
}

func (dto RecipeDto) GetModel() models.Recipe {
	return models.Recipe{
		
	}
}
