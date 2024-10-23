package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodRepositoryInterface interface {
	GetAll(userCode string, minimumList bool) ([]models.Food, error)
	GetByCode(foodCode primitive.ObjectID, userCode string) (models.Food, error)
	Create(newFood models.Food) (*mongo.InsertOneResult, error)
	Update(updateFood models.Food) (*mongo.UpdateResult, error)
	Delete(foodcode primitive.ObjectID) (*mongo.DeleteResult, error)
}

type FoodRepository struct {
	db DB
}

func NewFoodRepository(db DB) *FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (foodRepository FoodRepository) GetAll(userCode string, minimumList bool) ([]models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"user_code": userCode}

	if minimumList {
		filter = bson.M{
			"$expr": bson.M{
				"$lt": bson.A{"$current_quantity", "$minimum_quantity"},
			},
			"user_code": userCode,
		}
	}

	cursor, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var foods []models.Food
	cursor.All(context.TODO(), &foods)

	if len(foods) == 0 {
		err = errors.New("nocontent")
		return nil, err
	}
	return foods, nil
}

func (foodRepository FoodRepository) GetByCode(foodCode primitive.ObjectID, userCode string) (models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")
	
	filter := bson.M{
		"_id": foodCode,
		"user_code":   userCode,
	}
	data := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").FindOne(context.TODO(), filter)
	var food models.Food
	err := data.Decode(&food)

	if err == mongo.ErrNoDocuments {
		err = errors.New("could not find the food with the given code ")
	}
	return food, err
}

func (foodRepository FoodRepository) Create(food models.Food) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").InsertOne(context.TODO(), food)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (foodRepository FoodRepository) Update(food models.Food) (*mongo.UpdateResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"food_code": food.Code}
	update := bson.M{
		"$set": food, //actualiza solo los campos que esten en la variable food
	}
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (foodRepository FoodRepository) Delete(foodCode primitive.ObjectID) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"_id": foodCode}
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").DeleteOne(context.TODO(), filter)
	if res.DeletedCount == 0 {
		err = errors.New("notfound")
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

