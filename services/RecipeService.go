package services

import (
	"Status418/dto"
	"Status418/enums"
)

type RecipeServiceInterface interface {
	Create(newRecipe dto.RecipeDTO) error
	Delete(id int) error
	Update(updateRecipe dto.RecipeDTO) error
	GetByMoment(moment enums.Moment) ([]dto.RecipeDTO, error)
	GetByType(types enums.FoodType) ([]dto.RecipeDTO, error)
	GetAll() ([]dto.RecipeDTO, error)
}

