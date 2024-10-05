package dto

import "Status418/enums"

type FoodDto struct {
	Type      enums.FoodType 
	Moments    []enums.Moment 
	Name      string 
	UnitPrice float64 
	CurrentQuantity int
	MinimumQuantity int 
}

func NewFoodDTO(ftype enums.FoodType, moment []enums.Moment, name string, unitPrice float64, currentQuantity int, minimumQuantity int, userId string) *FoodDto {
	return &FoodDto{
		Type: ftype,
		Moments: moment,
		Name: name,
		UnitPrice: unitPrice,
		CurrentQuantity: currentQuantity,
		MinimumQuantity: minimumQuantity,
	}
}
