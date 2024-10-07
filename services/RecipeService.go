package services

import (
	"Status418/dto"
	"Status418/enums"
	"Status418/repositories"
)

type RecipeServiceInterface interface {
	GetByMoment(moment enums.Moment) ([]dto.RecipeDto, error)
	GetByType(types enums.FoodType) ([]dto.RecipeDto, error)
    GetAll() ([]dto.RecipeDto, error)
	Create(newRecipe dto.RecipeDto) error
	Delete(id string) error
	Update(updateRecipe dto.RecipeDto) error
}

type RecipeService struct{
	rr repositories.RecipeRepository
}

func NewRecipeService(rr repositories.RecipeRepository) *RecipeService{
	return &RecipeService{
		rr: rr,
	}
}

func (rr)
//dale mateo porfa