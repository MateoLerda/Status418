package services

import (
	"Status418/dto"
	"Status418/repositories"
)

type FoodServiceInterface interface {
	GetAll() (*[]dto.FoodDto, error)
	GetByCode(id int) (*dto.FoodDto, error)
	Create(newFood dto.FoodDto) error
	Update(updateFood dto.FoodDto) error
	Delete(id int) error
}

type FoodService struct {
	foodRepository repositories.FoodRepositoryInterface
}

func (ps *PurchaseService) GetFoodWithQuantityLessThanMinimum(userId string) (*[]dto.FoodDto, error) {
	var foodsDTO []dto.FoodDto
	foods, err := ps.pr.GetFoodWithQuantityLessThanMinimum(userId)
	if err != nil {
		return nil, err
	}
	for _, food := range *foods {
		foodDTO := dto.FoodDto{
			Type:            food.Type,
			Moments:         food.Moments,
			Name:            food.Name,
			UnitPrice:       food.UnitPrice,
			CurrentQuantity: food.CurrentQuantity,
			MinimumQuantity: food.MinimumQuantity,
		}
		foodsDTO = append(foodsDTO, foodDTO)
	}
	return &foodsDTO, nil
}

