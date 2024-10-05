package models

import (
	"Status418/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	Id_Recipe    primitive.ObjectID `bson:"id_recipe"`
	Name         string             `bson:"name"`
	Ingredients  []Food             `bson:"ingredients"`
	Moment       enums.Moment       `bson:"recipe_moment"`
	Description  string             `bson:"description"`
	CreationDate string             `bson:"creation_date"`
	UpdateDate   string             `bson:"update_date"`
	UserId       primitive.ObjectID `bson:"user_id"`
}

//AGREGAR EN LA BASE DE DATOS EL CAMPO DE MOMENTO
