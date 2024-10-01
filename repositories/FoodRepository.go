package repositories

import (
	"Status418/models"
	
)

type FoodRepositoryInterface interface {
	GetAll() ([]models.Food, error)
	GetByCode(id int) (models.Food, error)
	Create(models.Food) error
	Update(models.Food) error
	Delete(id int) error
}

