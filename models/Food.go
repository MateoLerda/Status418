package models

import "Status418/enums"

type Food struct {
	Code      int `bson:"food_code,omitempty"`
	Type      enums.FoodType `bson:"type"`
	Moment    enums.Moment `bson:"moment"`
	Name      string `bson:"name"`
	UnitPrice float64 `bson:"price"`
	CurrentQuantity int`bson:"current_quantity"`
	MinimumQuantity int `bson:"minimum_quantity"`
	CreationDate string `bson:"creation_date"`
	UpdateDate string `bson:"update_"`
	UserId int `bson:"user_id"`
}
