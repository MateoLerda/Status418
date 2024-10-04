package models

import (
	"Status418/enums"
)

type Food struct {
	Code            int            `bson:"food_code,omitempty"`
	Type            enums.FoodType `bson:"type"`
	Moments         []enums.Moment `bson:"moments"`
	Name            string         `bson:"name"`
	UnitPrice       float64        `bson:"price"`
	CurrentQuantity int            `bson:"current_quantity"`
	MinimumQuantity int            `bson:"minimum_quantity"`
	CreationDate    string         `bson:"creation_date"`
	UpdateDate      string         `bson:"update_"`
	UserId          int            `bson:"user_id"`
}
//CAMBIAR MOMENTO COMO ENUMERADOR Y COLOCAR UN SLICE DE MOMENTOS