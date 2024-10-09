package repositories

import (
	"Status418/models"
	"context"
	"errors"
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


func (pr PurchaseRepository) Create(purchase models.Purchase) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	// la compra ya viene completa desde el service, ahí llamamos al otro método de food y la creamos
	res, err := pr.db.GetClient().Database(DBNAME).Collection("Purchases").InsertOne(context.TODO(), purchase)
	if err != nil {
		err = errors.New("failed to create purchase")
		return nil, err
	}
	return res, nil
}

