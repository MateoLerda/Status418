package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PurchaseServiceInterface interface {
	Create(dto.PurchaseDTO) (*mongo.InsertOneResult, error)
}

type PurchaseService struct {
	pr repositories.PurchaseRepository
}

func NewPurchaseService(pr repositories.PurchaseRepository) *PurchaseService {
	return &PurchaseService{
		pr: pr,
	}
}


func (ps *PurchaseService) Create(purchaseDTO *dto.PurchaseDTO) (*mongo.InsertOneResult, error) {
	foods, err := ps.pr.GetFoodWithQuantityLessThanMinimum()
	if err != nil {
		return nil, err
	}
	var purchase models.Purchase
	purchase.UserId = purchaseDTO.UserId
	purchase.PurchaseDate = time.Now()
	for _, food := range *foods {
		purchase.TotalCost += food.UnitPrice * (float64)(food.MinimumQuantity-food.CurrentQuantity)
		purchase.Foods = append(purchase.Foods, dto.PurchaseQuantity{
			FoodCode: food.Code,
			Quantity: food.MinimumQuantity - food.CurrentQuantity,
		})
	}
	res, err := ps.pr.Create(&purchase)
	if err != nil {
		return nil, err
	}
	return res, nil
}
