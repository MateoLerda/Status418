package models

import "Status418/dto"

type Purchase struct {
	Id_Purchase  	  int     `bson:"id_purchase"`
	UserId    int     `bson:"user_id"`
	PurchaseDate string `bson:"purchase_date"`
	TotalCost float64 `bson:"total_cost"`
	//PROPIEDAD FOODS[] HAY QUE VER DE QUE TIPO ES YA QUE NECESITA LA CANTIDAD DE CADA ALIMENTO QUE SE COMPRA 
	Foods[] dto.PurchaseQuantity `bson:"foods"`
}