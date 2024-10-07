package dto

import (
	"Status418/enums"
	"Status418/models"
	"Status418/utils"
)

type FoodDto struct {
	Code            string
	Type            enums.FoodType
	Moments         []enums.Moment
	Name            string
	UnitPrice       float64
	CurrentQuantity int
	MinimumQuantity int
}

func NewFoodDto(model models.Food) *FoodDto {
	return &FoodDto{
		Code:            utils.GetStringIDFromObjectID(model.Code),
		Type:            model.Type,
		Moments:         model.Moments,
		Name:            model.Name,
		UnitPrice:       model.UnitPrice,
		CurrentQuantity: model.CurrentQuantity,
		MinimumQuantity: model.MinimumQuantity,
	}
}

func (dto FoodDto) GetModel() models.Food {
	return models.Food{
		Code:            utils.GetObjectIDFromStringID(dto.Code),
		Type:            dto.Type,
		Moments:         dto.Moments,
		Name:            dto.Name,
		UnitPrice:       dto.UnitPrice,
		CurrentQuantity: dto.CurrentQuantity,
		MinimumQuantity: dto.MinimumQuantity,
	}
}
