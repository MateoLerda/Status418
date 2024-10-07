package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeServiceInterface interface {
    GetAll(userId string, filters models.Filter) (*[]dto.RecipeDto, error)
	Create(newRecipe dto.RecipeDto) (*mongo.InsertOneResult ,error)
	Delete(id string) (*mongo.DeleteResult, error)
	Update(updateRecipe dto.RecipeDto) (*mongo.UpdateResult, error)
}

type RecipeService struct{
	rr repositories.RecipeRepositoryInterface
}

func NewRecipeService(rr repositories.RecipeRepositoryInterface) *RecipeService{
	return &RecipeService{
		rr: rr,
	}
}

func (rs *RecipeService) GetAll(userId string, filters models.Filter) (*[]dto.RecipeDto, error){
	var recipesDto []dto.RecipeDto
	recipes, err := rs.rr.GetAll(userId, filters)
	if err != nil{
		return nil, err
	}
	for _,recipe := range recipes {
		recipeDto := dto.NewRecipeDto(recipe)
		recipesDto = append(recipesDto, *recipeDto)
	}
	return &recipesDto, nil
}

func (rs *RecipeService) Create(newRecipe dto.RecipeDto) (*mongo.InsertOneResult ,error){
	recipe := newRecipe.GetModel()
	recipe.CreationDate = time.Now().String()
	res, err := rs.rr.Create(recipe)
	if err != nil{
		return nil, err
	}
	return res, nil
}

func (rs *RecipeService) Delete(id string) (*mongo.DeleteResult ,error){
	res, err := rs.rr.Delete(id)
	if err!= nil{
		return nil, err
	}
	 return res, nil
}

func (rs *RecipeService) Update(updateRecipe dto.RecipeDto) (*mongo.UpdateResult, error){
	recipe := updateRecipe.GetModel()
	recipe.UpdateDate = time.Now().String()
	res, err := rs.rr.Update(recipe)
	if err != nil{
		return nil, err
	}
	return res, nil
}