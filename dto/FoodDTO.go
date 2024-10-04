package dto

import "Status418/enums"

type FoodDTO struct {
	Type      enums.FoodType 
	Moments    []enums.Moment 
	Name      string 
	UnitPrice float64 
	CurrentQuantity int
	MinimumQuantity int 
	UserId int 
}

func NewFoodDTO(ftype enums.FoodType, moment []enums.Moment, name string, unitprice float64, currentquantity int, minimumquantity int, userid int) *FoodDTO {
	return &FoodDTO {
		Type: ftype,
		Moments: moment,
		Name: name,
		UnitPrice: unitprice,
		CurrentQuantity: currentquantity,
		MinimumQuantity: minimumquantity,
		UserId: userid,
	}
}
