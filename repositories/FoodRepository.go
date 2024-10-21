package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type FoodRepositoryInterface interface {
	GetAll(userCode string, minimumList bool) ([]models.Food, error)
	GetByCode(foodCode string, userCode string) (models.Food, error)
	Create(models.Food) (*mongo.InsertOneResult, error)
	Update(models.Food) (*mongo.UpdateResult, error)
	Delete(foodCode string) (*mongo.DeleteResult, error)
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
		err = errors.New("Not found any foods")
		return nil, err
	}
	return foods, nil
}

func (foodRepository FoodRepository) GetByCode(foodCode string, userCode string) (models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{
		"food_code": foodCode,
		"user_id":   userCode,
	}
	data := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").FindOne(context.TODO(), filter)
	var food models.Food
	err := data.Decode(&food)

	if err == mongo.ErrNoDocuments {
		err = errors.New("Could not find the food with the given code " + foodCode)
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
		"$set": food, //actualiza sólo los campos que estén en la variable food
	}
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (foodRepository FoodRepository) Delete(foodCode string) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"food_code": foodCode}
	res, err := foodRepository.db.GetClient().Database(DBNAME).Collection("Foods").DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
