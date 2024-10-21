package services

import (
	"Status418/dto"
	"Status418/enums"
	"Status418/models"
	"Status418/repositories"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeServiceInterface interface {
    GetAll(userCode string, filters dto.FiltersDto) (*[]dto.RecipeDto, error)
	Create(newRecipe dto.RecipeDto) (*mongo.InsertOneResult ,error)
	Delete(recipeId string) (*mongo.DeleteResult, error)
	Update(updateRecipe dto.RecipeDto) (*mongo.UpdateResult, error)
}

type RecipeService struct{
	recipeRepository repositories.RecipeRepositoryInterface
}

func NewRecipeService(recipeRepository repositories.RecipeRepositoryInterface) *RecipeService{
	return &RecipeService{
		recipeRepository: recipeRepository,
	}
}

func (recipeService *RecipeService) GetAll(userCode string, filters dto.FiltersDto) (*[]dto.RecipeDto, error){
	var recipesDto []dto.RecipeDto
	recipes, err := recipeService.recipeRepository.GetAll(userCode, filters.GetModel())
	if err != nil{
		return nil, err
	}
	for _,recipe := range recipes {
		recipeDto := dto.NewRecipeDto(recipe)
		recipesDto = append(recipesDto, *recipeDto)
	}
	return &recipesDto, nil
}

func (recipeService *RecipeService) Create(newRecipe dto.RecipeDto) (*mongo.InsertOneResult ,error){
	recipe := newRecipe.GetModel()
	recipe.CreationDate = time.Now().String()
	validation := find(recipe.Ingredients, recipe.Moment)
	if(!validation){
		return nil, errors.New("The food moment doesn´t match with the recipe moment")
	}
	res, err := recipeService.recipeRepository.Create(recipe)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func find(ingredients []models.Food, recipeMoment enums.Moment) bool {
	for _,food := range ingredients{
		momentValidation := false
		for _, moment := range food.Moments{
			if(moment == recipeMoment){
				momentValidation= true
				break;
			}
		}
		if(!momentValidation) {
			return false
		}
	}
	return true
}

func (recipeService *RecipeService) Delete(recipeId string) (*mongo.DeleteResult ,error){
	res, err := recipeService.recipeRepository.Delete(recipeId)
	if err!= nil{
		return nil, err
	}
	 return res, nil
}

func (recipeService *RecipeService) Update(updateRecipe dto.RecipeDto) (*mongo.UpdateResult, error){
	recipe := updateRecipe.GetModel()
	recipe.UpdateDate = time.Now().String()
	res, err := recipeService.recipeRepository.Update(recipe)
	if err != nil{
		return nil, err
	}
	return res, nil
}