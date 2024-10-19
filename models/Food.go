package models

import (
	"Status418/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	Code            primitive.ObjectID `bson:"food_code,omitempty"`
	Type            enums.FoodType     `bson:"type"`
	Moments         []enums.Moment     `bson:"moments"`
	Name            string             `bson:"name"`
	UnitPrice       float64            `bson:"price"`
	CurrentQuantity int                `bson:"current_quantity"`
	MinimumQuantity int                `bson:"minimum_quantity"`
	CreationDate    string             `bson:"creation_date"`
	UpdateDate      string             `bson:"update_"`
	UserCode        string             `bson:"user_code"`
}
