package dto

type RecipeDTO struct {
	Name string 
	Ingredients[] FoodDTO 
	Description string 
	UserId int 
}

func NewRecipeDTO(name string, ingredients []FoodDTO, description string, userid int) *RecipeDTO {
	return &RecipeDTO{
		Name: name,
		Ingredients: ingredients,
		Description: description,
		UserId: userid,
	}
}