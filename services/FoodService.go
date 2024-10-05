package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
	"time"
)

type FoodServiceInterface interface {
	GetAll() (*[]dto.FoodDto, error)
	GetByCode(id int) (*dto.FoodDto, error)
	Create(newFood dto.FoodDto) error
	Update(updateFood dto.FoodDto) error
	Delete(id int) error
}

type FoodService struct {
	fr repositories.FoodRepository
}

func NewFoodService(fr repositories.FoodRepository) *FoodService {
	return &FoodService{
		fr: fr,
	}
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

func (fs *FoodService) GetAll(userId string) (*[]dto.FoodDto, error) {
	var foodsDTO []dto.FoodDto
	foods, err := fs.fr.GetAll(userId) //CHEQUEAR, NO CREO QUE ESTÉ CORRECTO
	if err != nil {
		return nil, err
	}
	for _, food := range *foods {
		foodDTO := ChangeFromFoodModelToDto(&food)
		foodsDTO = append(foodsDTO, *foodDTO)
	}
	return &foodsDTO, nil
}

func (fs *FoodService) GetByCode(code string, userId string) (*dto.FoodDto, error){
	food, err := fs.fr.GetByCode(code, userId)
	if err!=nil {
		return nil, err
	}
	foodDto:= ChangeFromFoodModelToDto(food)
	return foodDto, nil
}

func (fs *FoodService) Create(foodDto *dto.FoodDto) error{
	food:= ChangeFromFoodDtoToModel(foodDto)
	food.CreationDate = time.Now().String()
	_, err := fs.fr.Create(food)
	if err != nil{
		return err
	}
	return nil
}

func (fs *FoodService) Update(foodDto *dto.FoodDto) error{
	food:= ChangeFromFoodDtoToModel(foodDto)
	food.UpdateDate = time.Now().String()
	_, err := fs.fr.Update(food)
	if err != nil{
		return err
	}
	return nil
}

func (fs *FoodService) Delete(code string) error{
	_, err := fs.fr.Delete(code)
	if err != nil{
		return err
	}
	return nil
}


func ChangeFromFoodDtoToModel(foodDTO *dto.FoodDto) *models.Food {
	food := models.Food{
		Type:			foodDTO.Type,
		Moments:		foodDTO.Moments,
		Name:			foodDTO.Name,
		UnitPrice:		foodDTO.UnitPrice,
		CurrentQuantity:foodDTO.CurrentQuantity,
		MinimumQuantity:foodDTO.MinimumQuantity,
		CreationDate:	time.Now().String(),
	}
	return &food
}

func ChangeFromFoodModelToDto(food *models.Food) *dto.FoodDto {
	foodDTO := dto.FoodDto{
		Type:            food.Type,
		Moments:         food.Moments,
		Name:            food.Name,
		UnitPrice:       food.UnitPrice,
		CurrentQuantity: food.CurrentQuantity,
		MinimumQuantity: food.MinimumQuantity,
	}
	return &foodDTO
}

