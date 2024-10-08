package dto

import (
	"Status418/enums"
	"Status418/models"
)

type FiltersDto struct {
	Aproximation string 
	Moment enums.Moment
	Type   enums.FoodType
}

func (dto FiltersDto) GetModel() models.Filter {
	return models.Filter{
		Aproximation: dto.Aproximation,
		Moment: dto.Moment,
		Type: dto.Type,
	}
}