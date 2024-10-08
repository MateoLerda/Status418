package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)

type FoodRepositoryInterface interface {
	GetAll(userId string) ([]models.Food, error)
	GetByCode(code string, userId string) (models.Food, error)
	Create(models.Food) (*mongo.InsertOneResult, error)
	Update(models.Food) (*mongo.UpdateResult, error)
	Delete(code string) (*mongo.DeleteResult, error)
}

type FoodRepository struct {
	db DB
}

func NewFoodRepository(db DB) *FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (fr FoodRepository) GetAll(userId string) ([]models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter:= bson.M{
		"user_id": userId,
	}
	cursor, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filter)

	if err != nil {
		err = errors.New("internal")
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

func (fr FoodRepository) GetByCode(code string, userId string) (models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{
		"food_code": code,
		"user_id":   userId,
	}
	data := fr.db.GetClient().Database(DBNAME).Collection("Foods").FindOne(context.TODO(), filter)
	var food models.Food
	err := data.Decode(&food)
	if err != nil {
		err = errors.New("internal")
	}
	if err == mongo.ErrNoDocuments {
		err = errors.New("notfound")
	}
	return food, err
}

func (fr FoodRepository) Create(food models.Food) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	res, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").InsertOne(context.TODO(), food)

	if err != nil {
		err = errors.New("failed to create food")
		return nil, err
	}

	return res, nil
}

func (fr FoodRepository) Update(food models.Food) (*mongo.UpdateResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"food_code": food.Code}
	update := bson.M{
		"$set": food, //actualiza sólo los campos que estén en la variable food
	}
	res, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		err = errors.New("internal")
		return nil, err
	}

	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		err = errors.New("notfound")
		return nil, err
	}

	return res, nil
}

func (fr FoodRepository) Delete(code string) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"food_code": code}
	res, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").DeleteOne(context.TODO(), filter)
	if err != nil {
		err = errors.New("internal")
		return nil, err
	}

	if res.DeletedCount == 0 {
		err = errors.New("notfound")
		return nil,err
	}
	return res, nil
}
