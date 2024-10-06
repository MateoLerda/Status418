package handlers

import "Status418/services"

type RecipeHandler struct {
	rs services.RecipeServiceInterface
}

func NewRecipeHandler(rs services.RecipeServiceInterface) *RecipeHandler {
	return &RecipeHandler{
		rs : rs,
	}
}

//IMPLEMENTAR LOS MÉTODOS DE LA INTERFACE RecipeServiceInterface