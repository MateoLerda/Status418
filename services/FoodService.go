package services

import (
	"Status418/dto"
	"Status418/repositories"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodServiceInterface interface {
	GetAll(userId string) (*[]dto.FoodDto, error)
	GetByCode(code string, userId string) (*dto.FoodDto, error)
	Create(newFood dto.FoodDto) (*mongo.InsertOneResult , error)
	Update(updateFood dto.FoodDto) (*mongo.UpdateResult , error)
	Delete(code string) (*mongo.DeleteResult , error)
}

type FoodService struct {
	fr repositories.FoodRepositoryInterface
}

func NewFoodService(fr repositories.FoodRepositoryInterface) *FoodService {
	return &FoodService{
		fr: fr,
	}
}

func (fs *FoodService) GetAll(userId string) (*[]dto.FoodDto, error) {
	var foodsDTO []dto.FoodDto
	foods, err := fs.fr.GetAll(userId) 
	if err != nil {
		return nil, err
	}
	for _, food := range foods {
		foodDTO := dto.NewFoodDto(food)
		foodsDTO = append(foodsDTO, *foodDTO)
	}
	return &foodsDTO, nil
}

func (fs *FoodService) GetByCode(code string, userId string) (*dto.FoodDto, error){
	food, err := fs.fr.GetByCode(code, userId)
	if err!=nil {
		return nil, err
	}
	foodDto:= dto.NewFoodDto(food)
	return foodDto, nil
}

func (fs *FoodService) Create(foodDto dto.FoodDto) (*mongo.InsertOneResult , error) {
	food := foodDto.GetModel()
	food.CreationDate = time.Now().String()
	res, err := fs.fr.Create(food)
	if err != nil{
		return nil, err
	}
	return res, nil
}

func (fs *FoodService) Update(foodDto dto.FoodDto) (*mongo.UpdateResult , error) {
	food:= foodDto.GetModel()
	food.UpdateDate = time.Now().String()
	res, err := fs.fr.Update(food)
	if err != nil{
		return nil, err
	}
	return res, nil
}

func (fs *FoodService) Delete(code string) (*mongo.DeleteResult , error) {
	res, err := fs.fr.Delete(code)
	if err != nil{
		return nil, err
	}
	return res, nil
}



