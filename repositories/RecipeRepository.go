package repositories

import (
	"Status418/models"
	"Status418/utils"
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeRepositoryInterface interface {
	Create(recipe models.Recipe) (*mongo.InsertOneResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	Update(recipe models.Recipe) (*mongo.UpdateResult, error)
	GetAll(userId string, filters models.Filter) ([]models.Recipe, error)
}

type RecipeRepository struct {
	db DB
}

func NewRecipeRepository(db DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

func (rr RecipeRepository) Create(recipe models.Recipe) (*mongo.InsertOneResult, error) {
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
	idMongo := utils.GetObjectIDFromStringID(id)
	filter := bson.M{"id_recipe": idMongo}
	res, err := rr.db.GetClient().Database(DBNAME).Collection("Recipes").DeleteOne(context.TODO(), filter)
	if err != nil {
		err = errors.New("failed to delete recipe")
		return res, err
	}
	return res, nil
}

func (rr RecipeRepository) Update(recipe models.Recipe) (*mongo.UpdateResult, error) {
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
} 


func (rr RecipeRepository) GetAll(userId string, filters models.Filter) ([]models.Recipe, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{
		"name": bson.M{
			"$regex": filters.Aproximation,
			"$options": "i",
		},
		"ingredients": bson.M{
			"$elemMatch": bson.M{
				"type": filters.Type,
			},
		},
		"recipe_moment": filters.Moment, 
		"user_id": userId,
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

	return recipes, nil
}
