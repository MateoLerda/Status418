package models

import (
	"Status418/enums"
)

type Filter struct {
	Aproximation string 
	Moment enums.Moment
	Type   enums.FoodType
	All 	bool
}
