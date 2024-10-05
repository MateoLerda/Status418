package services

import (
	"Status418/dto"
	"Status418/enums"
)

type RecipeServiceInterface interface {
	Create(newRecipe dto.RecipeDto) error
	Delete(id int) error
	Update(updateRecipe dto.RecipeDto) error
	GetByMoment(moment enums.Moment) ([]dto.RecipeDto, error)
	GetByType(types enums.FoodType) ([]dto.RecipeDto, error)
	GetAll() ([]dto.RecipeDto, error)
}

