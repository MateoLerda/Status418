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
	GetAll() (*[]models.Food, error)
	GetByCode(id int) (*models.Food, error)
	Create(models.Food) (*mongo.InsertOneResult, error)
	Update(models.Food) (*mongo.UpdateResult, error)
	Delete(id int) (*mongo.DeleteResult, error)
	GetFoodWithQuantityLessThanMinimum() (*[]models.Food, error)
}

type FoodRepository struct {
	db DB
}

func NewFoodRepository(db DB) *FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (fr FoodRepository) GetAll() (*[]models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{}

	cursor, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filter)

	if err != nil {
		err = errors.New("failed to get foods")
		return nil, err
	}

	var foods []models.Food
	err = cursor.All(context.TODO(), &foods)

	if err != nil {
		err = errors.New("failed to parse food documents")
		return nil, err
	}
	return &foods, nil
}

func (fr FoodRepository) GetByCode(code int) (*models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{"food_code": code}
	data := fr.db.GetClient().Database(DBNAME).Collection("Foods").FindOne(context.TODO(), filter)

	if data == nil {

	}

	var food models.Food
	err := data.Decode(&food)
	if err != nil {
		err = errors.New("failed to get food")
		return nil, err
	}
	return &food, nil

}

func (fr FoodRepository) Create(food *models.Food) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	res, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").InsertOne(context.TODO(), food)

	if err != nil {
		err = errors.New("failed to create food")
		return nil, err
	}

	return res, nil
}

func (fr FoodRepository) Update(food *models.Food) (*mongo.UpdateResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"food_code": food.Code}
	update := bson.M{
		"$set": food, //actualiza sólo los campos que estén en la variable food
	}
	res, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		err = errors.New("failed to update food")
		return nil, err
	}
	return res, nil
}

func (fr FoodRepository) Delete(code int) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"food_code": code}
	res, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").DeleteOne(context.TODO(), filter)
	if err != nil {
		err = errors.New("failed to delete food")
		return nil, err
	}
	return res, nil
}

func (pr PurchaseRepository) GetFoodWithQuantityLessThanMinimum() (*[]models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{
		"$expr": bson.M{
			"$lt": bson.A{"$current_quantity", "$minimum_quantity"},
		},
	}
	cursor, err := pr.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filter)

	if err != nil {
		err = errors.New("failed to get foods with quantity less than minimum")
		return nil, err
	}

	var foods []models.Food
	err = cursor.All(context.TODO(), &foods)

	if err != nil {
		err = errors.New("failed to parse foods with quantity less than minimum documents")
		return nil, err
	}
	return &foods, nil
}
