package repositories

import (
	"Status418/models"
	"Status418/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type RecipeRepositoryInterface interface {
	Create(recipe models.Recipe) (*mongo.InsertOneResult, error)
	Delete(recipeId string) (*mongo.DeleteResult, error)
	Update(recipe models.Recipe) (*mongo.UpdateResult, error)
	GetAll(userCode string, filters models.Filter) ([]models.Recipe, error)
	GetByCode(userCode string, recipeId string) (models.Recipe, error)
}

type RecipeRepository struct {
	db DB
}

func NewRecipeRepository(db DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

func (recipeRepository RecipeRepository) Create(recipe models.Recipe) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	res, err := recipeRepository.db.GetClient().Database(DBNAME).Collection("Recipes").InsertOne(context.TODO(), recipe)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (recipeRepository RecipeRepository) Delete(recipeId string) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	mongoRecipeId := utils.GetObjectIDFromStringID(recipeId)
	filter := bson.M{"id_recipe": mongoRecipeId}
	res, err := recipeRepository.db.GetClient().Database(DBNAME).Collection("Recipes").DeleteOne(context.TODO(), filter)
	if err != nil {
		err = errors.New("internal")
		return nil, err
	}
	if res.DeletedCount == 0 {
		err = errors.New("notfound")
		return nil, err
	}
	return res, nil
}

func (recipeRepository RecipeRepository) Update(recipe models.Recipe) (*mongo.UpdateResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{
		"id_recipe": recipe.Id,
	}
	update := bson.M{
		"$set": recipe,
	}
	res, err := recipeRepository.db.GetClient().Database(DBNAME).Collection("Recipes").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		err = errors.New("internal")
		return res, err
	}

	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		err = errors.New("notfound")
		return nil, err
	}
	return res, nil
}

func (recipeRepository RecipeRepository) GetAll(userCode string, filters models.Filter) ([]models.Recipe, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{
		"recipe_name": bson.M{
			"$regex":   filters.Aproximation,
			"$options": "i",
		},
		"recipe_moment": filters.Moment,
		"user_code":     userCode,
	}

	data, err := recipeRepository.db.GetClient().Database(DBNAME).Collection("Recipes").Find(context.TODO(), filter)
	if err != nil {
		err = errors.New("internal")
		return nil, err
	}
	var recipes []models.Recipe
	data.All(context.TODO(), &recipes)
	if len(recipes) == 0 {
		err = errors.New("nocontent")
		return nil, err
	}

	return recipes, nil
}

func (recipeRepository RecipeRepository) GetByCode(userCode string, recipeId string) (models.Recipe, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{
		"id_recipe": recipeId,
		"user_code": userCode,
	}
	data := recipeRepository.db.GetClient().Database(DBNAME).Collection("Recipes").FindOne(context.TODO(), filter)
	var recipe models.Recipe
	err := data.Decode(&recipe)
	if err == mongo.ErrNoDocuments {
		err = errors.New("Couldn´t find the recipe with the id: " + recipeId)
	}

	return recipe, err
}
