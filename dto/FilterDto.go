package dto

import (
	"Status418/enums"
	"Status418/models"
)

type FiltersDto struct {
	Aproximation string 
	Moment string
	Type   string
}

func (dto FiltersDto) GetModel() models.Filter {
	return models.Filter{
		Aproximation: dto.Aproximation,
		Moment: enums.GetMomentEnum(dto.Moment),
		Type: enums.GetTypeEnum(dto.Type),
	}
}