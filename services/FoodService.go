package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
	"Status418/utils"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodServiceInterface interface {
	GetAll(userCode string, minimumList bool) (*[]dto.FoodDto, error)
	GetByCode(code string, userCode string) (*dto.FoodDto, error)
	Create(newFood dto.FoodDto, userCode string) (*mongo.InsertOneResult, error)
	Update(updateFood dto.FoodDto) (*mongo.UpdateResult, error)
	Delete(userCode string, foodcode string) (*mongo.DeleteResult, error)
}

type FoodService struct {
	foodRepository repositories.FoodRepositoryInterface
}

func NewFoodService(foodRepository repositories.FoodRepositoryInterface) *FoodService {
	return &FoodService{
		foodRepository: foodRepository,
	}
}

func (foodService *FoodService) GetAll(userCode string, minimumList bool) (*[]dto.FoodDto, error) {
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

func (foodService *FoodService) GetByCode(code string, userCode string) (*dto.FoodDto, error) {
	food, err := foodService.foodRepository.GetByCode(code, userCode)
	if err != nil {
		return nil, err
	}
	foodDto := dto.NewFoodDto(food)
	return foodDto, nil
}

func (foodService *FoodService) Create(foodDto dto.FoodDto, userCode string) (*mongo.InsertOneResult, error) {
	food := foodDto.GetModel()
	food.CreationDate = time.Now().String()
	food.UserCode = userCode
	res, err := foodService.foodRepository.Create(food)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (foodService *FoodService) Update(foodDto dto.FoodDto) (*mongo.UpdateResult, error) {
	food := foodDto.GetModel()
	food.UpdateDate = time.Now().String()
	res, err := foodService.foodRepository.Update(food)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (foodService *FoodService) Delete(userCode string, foodCode string) (*mongo.DeleteResult, error) {
	DB := repositories.NewMongoDB()
	recipeRepository := repositories.NewRecipeRepository(DB)
	recipes , _ := recipeRepository.GetAll(userCode, models.Filter{})
	for _,recipe := range recipes {
		for _, food := range recipe.Ingredients {
			if food.FoodCode == foodCode {
				recipeRepository.Delete(utils.GetStringIDFromObjectID(recipe.Id))
				break
			}
		}
	}
	
	res, err := foodService.foodRepository.Delete(foodCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}
