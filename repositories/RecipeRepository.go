package repositories

import (
	"Status418/enums"
	"Status418/models"
	"context"
	"errors"
	"os"
	"go.mongodb.org/mongo-driver/bson"
)

type RecipeRepositoryInterface interface {
	Create(newRecipe models.Recipe) error
	Delete(id int) error
	Update(updateRecipe models.Recipe) error
	GetByMoment(moment enums.Moment) ([]models.Recipe, error)
	GetByType(types enums.FoodType) ([]models.Recipe, error)
	GetAll() ([]models.Recipe, error)
}

type RecipeRepository struct {
	db DB
}

func NewRecipeRepository(db DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

func (rr RecipeRepository) Create(recipe *models.Recipe) error  {
	DBNAME:= os.Getenv("DB_NAME")
	_, err:= rr.db.GetClient().Database(DBNAME).Collection("Recipes").InsertOne(context.TODO(), recipe)
	if err!= nil{
		err = errors.New("failed to create recipe")
		return err
	}
	return nil
}

func (rr RecipeRepository) Delete(id int) error {
	DBNAME:= os.Getenv("DB_NAME")
	filtro:= bson.M{"id_recipe": id}
	_, err:= rr.db.GetClient().Database(DBNAME).Collection("Recipes").DeleteOne(context.TODO(),filtro)
	if err!= nil{
		err = errors.New("failed to delete recipe")
		return err
	}
	return nil
}

func (rr RecipeRepository) Update(recipe *models.Recipe) error {
	DBNAME := os.Getenv("DB_NAME")
	filtro := bson.M{"id_recipe": recipe.Id_Recipe}
	_, err := rr.db.GetClient().Database(DBNAME).Collection("Recipes").UpdateOne(context.TODO(),filtro,recipe)
	if err!= nil{
		err = errors.New("failed to update recipe")
		return err
	}
	return nil
}

func (rr RecipeRepository) GetByMoment(moment enums.Moment) (*[]models.Recipe, error) {
	return &[]models.Recipe{}, nil
}

func (rr RecipeRepository) GetByType(types enums.FoodType) (*[]models.Recipe, error) {
	
	return &[]models.Recipe{}, nil
}

func (rr RecipeRepository) GetAll() (*[]models.Recipe, error) {
	DBNAME := os.Getenv("DB_NAME")
	filtro := bson.M{}
	data, err := rr.db.GetClient().Database(DBNAME).Collection("Recipes").Find(context.TODO(),filtro)
	var recipes []models.Recipe
	err = data.All(context.TODO(), &recipes)
	if err != nil{
		err = errors.New("failed to get all recipes")
	}
	
	return &recipes, nil
}


