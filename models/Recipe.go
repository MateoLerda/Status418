package models

import (
	"Status418/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	Id           primitive.ObjectID `bson:"id_recipe"`
	Name         string             `bson:"recipe_name"`
	Ingredients  []FoodQuantity    	`bson:"ingredients"`
	Moment       enums.Moment       `bson:"recipe_moment"`
	Description  string             `bson:"description"`
	CreationDate string             `bson:"creation_date"`
	UpdateDate   string             `bson:"update_date"`
	UserCode     string			    `bson:"user_code"`
}

