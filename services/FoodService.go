package services

import (
	"Status418/dto"
	"Status418/repositories"
)

type FoodServiceInterface interface {
	GetAll() ([]dto.FoodDTO, error)
	GetByCode(id int) (dto.FoodDTO, error)
	Crete(newFood dto.FoodDTO) error
	Update(updateFood dto.FoodDTO) error
	Delete(id int) error
}

type FoodService struct{
	foodRepository repositories.FoodRepositoryInterface
}