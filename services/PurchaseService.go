package services

import (
	"Status418/models"
	"Status418/repositories"
	"Status418/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PurchaseServiceInterface interface {
	Create(userId string) (*mongo.InsertOneResult, error)
}

type PurchaseService struct {
	pr repositories.PurchaseRepository
}

func NewPurchaseService(pr repositories.PurchaseRepository) *PurchaseService {
	return &PurchaseService{
		pr: pr,
	}
}


func (ps *PurchaseService) Create(userId string) (*mongo.InsertOneResult, error) {
	foods, err := ps.pr.GetFoodWithQuantityLessThanMinimum(userId)
	
	if err != nil {
		return nil, err
	}
	var purchase models.Purchase

	purchase.PurchaseDate = time.Now()
	for _, food := range *foods {
		purchase.TotalCost += food.UnitPrice * (float64)(food.MinimumQuantity-food.CurrentQuantity)
		purchase.Foods = append(purchase.Foods, models.PurchaseQuantity{
			FoodCode: utils.GetStringIDFromObjectID(food.Code),
			Quantity: food.MinimumQuantity - food.CurrentQuantity,
		})
	}
	res, err := ps.pr.Create(&purchase)
	if err != nil {
		return nil, err
	}
	return res, nil
}
