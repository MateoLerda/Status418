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
	Cook(userCode string, recipeId string) (bool,error)
	CancelationCook(userCode string, recipeId string) (bool, error)
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
	if err != nil {
		return nil, err
	}
	for _ , recipe := range recipes {
		for _, FoodQuantity := range recipe.Ingredients {
			food, err := getFoodByCode(FoodQuantity.FoodCode, userCode)
			if(err != nil){
				return nil, err 
			}
			if (food.CurrentQuantity >= FoodQuantity.Quantity) {		
				recipeDto := dto.NewRecipeDto(recipe)
				recipesDto = append(recipesDto, *recipeDto)
			}
		}
	}
	return &recipesDto, nil
}

func (recipeService *RecipeService) Create(newRecipe dto.RecipeDto) (*mongo.InsertOneResult ,error){
	recipe := newRecipe.GetModel()
	recipe.CreationDate = time.Now().String()
	validation, err:= find(recipe.Ingredients, recipe.Moment, recipe.UserCode)
	if(err != nil){
		return nil, err
	}
	if(!validation){
		return nil, errors.New("The food moment doesn´t match with the recipe moment")
	}
	res, err := recipeService.recipeRepository.Create(recipe)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func find(ingredients []models.FoodQuantity, recipeMoment enums.Moment, userCode string) (bool, error) {
	for _,ingredient := range ingredients{
		momentValidation := false
		food, err:= getFoodByCode(ingredient.FoodCode,userCode)
		if(err!= nil){
			return false, err
		}
		for _, moment := range food.Moments {
			if(moment == recipeMoment){
				momentValidation= true
				break;
			}
		}
		if(!momentValidation) {
			return false,nil
		}
	}
	return true, nil
}

func getFoodByCode(FoodCode string, userCode string) (*models.Food, error){
	DB := repositories.NewMongoDB()
	foodRepository := repositories.NewFoodRepository(DB)
	food, err := foodRepository.GetByCode(FoodCode , userCode)
	if(err != nil){
		return nil, err 
	}
	return &food, nil
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

func (recipeService *RecipeService) Cook(userCode string, recipeId string) (bool,error){
	recipe, err:= recipeService.recipeRepository.GetByCode(userCode, recipeId) 
	if(err!= nil){
		return false, err
	} 
	DB := repositories.NewMongoDB()
	foodRepository := repositories.NewFoodRepository(DB)
	for _,foodQuantity := range recipe.Ingredients { 
		food, err := getFoodByCode(foodQuantity.FoodCode, userCode)
		if(err!= nil){
			return false, err
		}
		food.CurrentQuantity -= foodQuantity.Quantity
		_, err = foodRepository.Update(*food);
		if(err != nil){
			return false, err
		}
	}
	return true, nil
}

func (recipeService *RecipeService) CancelationCook(userCode string, recipeId string) (bool, error){
	recipe, err:= recipeService.recipeRepository.GetByCode(userCode, recipeId) 
	if(err!= nil){
		return false, err
	} 
	DB := repositories.NewMongoDB()
	foodRepository := repositories.NewFoodRepository(DB)
	for _,foodQuantity := range recipe.Ingredients { 
		food, err := getFoodByCode(foodQuantity.FoodCode, userCode)
		if(err!= nil){
			return false, err
		}
		food.CurrentQuantity += foodQuantity.Quantity
		_, err = foodRepository.Update(*food);
		if(err != nil){
			return false, err
		}
	}
	return true, nil
}
