package models

import "Status418/enums"

type Recipe struct {
	Id_Recipe    int    `bson:"id_recipe"`
	Name         string `bson:"name"`
	Ingredients  []Food `bson:"ingredients"`
	Moment       enums.Moment `bson:"recipe_moment"`
	Description  string `bson:"description"`
	CreationDate string `bson:"creation_date"`
	UpdateDate   string `bson:"update_date"`
	UserId       int    `bson:"user_id"`
}
//AGREGAR EN LA BASE DE DATOS EL CAMPO DE MOMENTO