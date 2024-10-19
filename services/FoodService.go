package services

import (
	"Status418/dto"
	"Status418/repositories"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodServiceInterface interface {
	GetAll(userCode string) (*[]dto.FoodDto, error)
	GetByCode(code string, userCode string) (*dto.FoodDto, error)
	Create(newFood dto.FoodDto) (*mongo.InsertOneResult , error)
	Update(updateFood dto.FoodDto) (*mongo.UpdateResult , error)
	Delete(code string) (*mongo.DeleteResult , error)
}

type FoodService struct {
	foodRepository repositories.FoodRepositoryInterface
}

func NewFoodService(foodRepository repositories.FoodRepositoryInterface) *FoodService {
	return &FoodService{
		foodRepository: foodRepository,
	}
}

func (foodService *FoodService) GetAll(userCode string) (*[]dto.FoodDto, error) {
	var foodsDTO []dto.FoodDto
	foods, err := foodService.foodRepository.GetAll(userCode, false) 
	if err != nil {
		return nil, err
	}
	for _, food := range foods {
		foodDTO := dto.NewFoodDto(food)
		foodsDTO = append(foodsDTO, *foodDTO)
	}
	return &foodsDTO, nil
}

func (foodService *FoodService) GetByCode(code string, userCode string) (*dto.FoodDto, error){
	food, err := foodService.foodRepository.GetByCode(code, userCode)
	if err!=nil {
		return nil, err
	}
	foodDto:= dto.NewFoodDto(food)
	return foodDto, nil
}

func (foodService *FoodService) Create(foodDto dto.FoodDto) (*mongo.InsertOneResult , error) {
	food := foodDto.GetModel()
	food.CreationDate = time.Now().String()
	res, err := foodService.foodRepository.Create(food)
	if err != nil{
		return nil, err
	}
	return res, nil
}

func (foodService *FoodService) Update(foodDto dto.FoodDto) (*mongo.UpdateResult , error) {
	food:= foodDto.GetModel()
	food.UpdateDate = time.Now().String()
	res, err := foodService.foodRepository.Update(food)
	if err != nil{
		return nil, err
	}
	return res, nil
}

func (foodService *FoodService) Delete(code string) (*mongo.DeleteResult , error) {
	res, err := foodService.foodRepository.Delete(code)
	if err != nil{
		return nil, err
	}
	return res, nil
}



