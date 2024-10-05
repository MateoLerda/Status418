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

func NewFoodDTO(ftype enums.FoodType, moment []enums.Moment, name string, unitprice float64, currentquantity int, minimumquantity int, userId string) *FoodDto {
	return &FoodDto{
		Type: ftype,
		Moments: moment,
		Name: name,
		UnitPrice: unitprice,
		CurrentQuantity: currentquantity,
		MinimumQuantity: minimumquantity,
	}
}
