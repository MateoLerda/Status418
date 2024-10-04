package models

type Recipe struct {
	Id_Recipe    int    `bson:"id_recipe"`
	Name         string `bson:"name"`
	Ingredients  []Food `bson:"ingredients"`
	Description  string `bson:"description"`
	CreationDate string `bson:"creation_date"`
	UpdateDate   string `bson:"update_date"`
	UserId       int    `bson:"user_id"`
}
