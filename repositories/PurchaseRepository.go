package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"
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
	DBNAME := os.Getenv("DB_NAME")
	data, err := pr.GetFoodWithQuantityLessThanMinimum()
	if err != nil {
		err = errors.New("cannot get the list of foods with quantity less than minimum")
		return err
	}
	
	for _, purchase := range *data {

		_, err := pr.db.GetClient().Database(DBNAME).Collection("Purchases").InsertOne(context.TODO(), purchase)
		if err != nil {
			err = errors.New("failed to create purchase")
			return err
		}
	}
	return nil
}
