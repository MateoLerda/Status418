package repositories

import (
	"Status418/enums"
	"Status418/models"
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeRepositoryInterface interface {
	Create(recipe *models.Recipe) (*mongo.InsertOneResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	Update(recipe *models.Recipe) (*mongo.UpdateResult, error)
	GetAll(aproximation string, userId string) (*[]models.Recipe, error)
}

type RecipeRepository struct {
	db DB
}

func NewRecipeRepository(db DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

func (rr RecipeRepository) Create(recipe *models.Recipe) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	res, err := rr.db.GetClient().Database(DBNAME).Collection("Recipes").InsertOne(context.TODO(), recipe)
	if err != nil {
		err = errors.New("failed to create recipe")
		return res, err
	}
	return res, nil
}

func (rr RecipeRepository) Delete(id string) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"id_recipe": id}
	res, err := rr.db.GetClient().Database(DBNAME).Collection("Recipes").DeleteOne(context.TODO(), filter)
	if err != nil {
		err = errors.New("failed to delete recipe")
		return res, err
	}
	return res, nil
}

func (rr RecipeRepository) Update(recipe *models.Recipe) (*mongo.UpdateResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{
		"id_recipe": recipe.Id,
	}
	update := bson.M {
		"$set": recipe,
	}
	res, err := rr.db.GetClient().Database(DBNAME).Collection("Recipes").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		err = errors.New("failed to update recipe")
		return res, err
	}
	return res, nil
} //CHEQUEAR ESTE UPDATE POR SI NO FUNCIONA Y SETEA LOS VALORES VACIOS EN LA BASE DE DATOS


func (rr RecipeRepository) GetAll(userId string, aproximation string,moment enums.Moment) (*[]models.Recipe, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{}
	if aproximation != "" {
		filter = bson.M{
			"name": bson.M{
				"$regex":   aproximation,
				"$options": "i",
			},
			"recipe_moment": moment, 
			"user_id": userId,
		}
	}
	data, err := rr.db.GetClient().Database(DBNAME).Collection("Recipes").Find(context.TODO(), filter)
	if err != nil {
		err = errors.New("failed to get all recipes")
		return nil, err
	}
	var recipes []models.Recipe
	err = data.All(context.TODO(), &recipes)
	if err != nil {
		err = errors.New("failed to parse all recipes")
		return nil, err
	}

	return &recipes, nil
}// A CHECKEAR COMPLETAMENTE 
