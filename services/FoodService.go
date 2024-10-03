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


	//MAPEO DE DATOS
	// var foods []models.Food
	// 	err = data.All(context.TODO(), &foods)

	// 	if err != nil {
	// 		err = errors.New("failed to get foods")
	// 		return nil, err
	// 	}

	// 	var foodDTOs []dto.FoodDTO
	// 	for _, food := range foods {
	// 		foodDTO := dto.FoodDTO{
	// 			Type:            food.Type,
	// 			Moment:          food.Moment,
	// 			Name:            food.Name,
	// 			UnitPrice:       food.UnitPrice,
	// 			CurrentQuantity: food.CurrentQuantity,
	// 			MinimumQuantity: food.MinimumQuantity,
	// 			UserId:          food.UserId,
	// 		}
	// 		foodDTOs = append(foodDTOs, foodDTO)
	// 	}