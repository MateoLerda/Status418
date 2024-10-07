package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PurchaseRepositoryInterface interface {
	Create(*models.Purchase) (*mongo.InsertOneResult, error)
	GetFoodWithQuantityLessThanMinimum() (*[]models.Food, error)
}

type PurchaseRepository struct {
	db DB
}

func NewPurchaseRepository(db DB) *PurchaseRepository {
	return &PurchaseRepository{
		db: db,
	}
}


func (pr PurchaseRepository) Create(purchase *models.Purchase) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")
	// la compra ya viene completa desde el service, ahí llamamos al otro método de food y la creamos
	res, err := pr.db.GetClient().Database(DBNAME).Collection("Purchases").InsertOne(context.TODO(), purchase)
	if err != nil {
		err = errors.New("failed to create purchase")
		return nil, err
	}
	return res, nil
}

func (pr PurchaseRepository) GetFoodWithQuantityLessThanMinimum(userId string) (*[]models.Food, error)  {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{
		"$expr": bson.M{
			"$lt": bson.A{"$current_quantity", "$minimum_quantity"},
		},
		"user_id": userId,
	}

	cursor, err := pr.db.GetClient().Database(DBNAME).Collection("Purchases").Find(context.TODO(), filter)

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

	return &foods,nil
}