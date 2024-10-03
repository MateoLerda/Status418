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
	GetAll() ([]models.Food, error)
	GetByCode(id int) (models.Food, error)
	Create(models.Food) error
	Update(models.Food) error
	Delete(id int) error
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
	filtro := bson.M{}

	cursor, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filtro)

	if err != nil {
		err = errors.New("failed to get foods")
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var foods []models.Food
	err = cursor.All(context.TODO(), &foods)

	if err != nil {
		err = errors.New("failed to parse food documents")
		return nil, err
	}
	return &foods, nil
}

func (fr FoodRepository) GetByCode(id int) (*models.Food, error) {
	DBNAME := os.Getenv("DB_NAME")

	filtro := bson.M{"food_code": id}
	data := fr.db.GetClient().Database(DBNAME).Collection("Foods").FindOne(context.TODO(), filtro)
	
	var food models.Food
	err := data.Decode(&food)
	if err != nil {
		err = errors.New("failed to get food")
		return nil, err
	}
	return &food, err

}

func (fr FoodRepository) Create(food *models.Food) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	res, err:= fr.db.GetClient().Database(DBNAME).Collection("Foods").InsertOne(context.TODO(), food)

	if err != nil {
		err = errors.New("failed to create food")
		return nil, err
	}
	
	return res, nil
}

func (fr FoodRepository) Update(food *models.Food) error {
	DBNAME := os.Getenv("DB_NAME")
	filtro := bson.M{"food_code": food.Code}
	_, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").UpdateOne(context.TODO(), filtro, food)
	if err != nil {
		err = errors.New("failed to update food")
		return err
	}
	return nil
}

func (fr FoodRepository) Delete(id int) error {
	DBNAME:= os.Getenv("DB_NAME")
	filtro := bson.M{"food_code": id}
	_, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").DeleteOne(context.TODO(), filtro)
	if err != nil {
		err = errors.New("failed to delete food")
		return err
	}
	return nil
}
