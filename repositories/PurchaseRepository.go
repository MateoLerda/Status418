package repositories

import (
	"Status418/models"
	"context"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
)

type PurchaseRepositoryInterface interface {
	Create(models.Purchase) (*mongo.InsertOneResult, error)
}

type PurchaseRepository struct {
	db DB
}

func NewPurchaseRepository(db DB) *PurchaseRepository {
	return &PurchaseRepository{
		db: db,
	}
}


func (purchaseRepository PurchaseRepository) Create(purchase models.Purchase) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")

	res, err := purchaseRepository.db.GetClient().Database(DBNAME).Collection("Purchases").InsertOne(context.TODO(), purchase)
	if err != nil {
		return nil, err
	}
	return res, nil
}

