package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

type PurchaseRepositoryInterface interface {
	GetFoodWithQuantityLessThanMinimum() ([]models.Food,error)
	Create(models.Purchase)error
}

type PurchaseRepository struct {
	db DB
}

func NewPurchaseRepository(db DB) *PurchaseRepository {
	return &PurchaseRepository{
		db: db,
	}
}

func (pr PurchaseRepository) GetFoodWithQuantityLessThanMinimum() (*[]models.Food,error) {
	envRes := godotenv.Load(".env")
	if envRes != nil {
		return nil, envRes
	}
	DBNAME := os.Getenv("DB_NAME")
	filtro:= bson.M{ 
		"$expr": bson.M{
			"$lt": bson.A{"$current_quantity","$minimum_quantity"},
		},
	}
	data, err := pr.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filtro)
	if err != nil {
		err = errors.New("failed to get foods with quantity less than minimum")
		return nil, err
	}
	var foods []models.Food
	err = data.All(context.TODO(), &foods)
	if err != nil {
		err = errors.New("failed to get foods with quantity less than minimum")
		return nil, err
	}
	return &foods, nil
}

func (pr PurchaseRepository) Create(purchase *models.Purchase) error {
	envRes := godotenv.Load(".env")
	if envRes != nil {
		return envRes
	}
	DBNAME := os.Getenv("DB_NAME")
	data, err := pr.GetFoodWithQuantityLessThanMinimum()
	if err != nil {
		err = errors.New("Cannot make the list of foods with quantity less than minimum")
		return err
	}
	var foodsLessThanMinimum = *data
	for _, food := range foodsLessThanMinimum {
		purchase.Food = food
		_, err := pr.db.GetClient().Database(DBNAME).Collection("Purchases").InsertOne(context.TODO(), purchase)
		if err != nil {
			err = errors.New("failed to create purchase")
			return err
		}
	}
	return nil
}
