package dto

type RecipeDto struct {
	Name string 
	Ingredients []FoodDto
	Description string 
	UserId string 
}

func NewRecipeDto(name string, ingredients []FoodDto, description string, userId string) *RecipeDto {
	return &RecipeDto{
		Name: name,
		Ingredients: ingredients,
		Description: description,
		UserId: userId,
	}
}