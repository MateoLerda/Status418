package repositories

import (
	"Status418/enums"
	"Status418/models"
)

type RecipeRepositotory interface {
	Create(newRecipe models.Recipe) error
	Delete(id int) error
	Update(updateRecipe models.Recipe) error
	GetByMoment(moment enums.Moment) ([]models.Recipe,error)
	GetByType(types enums.FoodType) ([]models.Recipe,error)
	GetAll()([]models.Recipe,error)
}