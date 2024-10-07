package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
	"Status418/utils"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type PurchaseServiceInterface interface {
	Create(userId string) (*mongo.InsertOneResult, error)
	GetFoodWithQuantityLessThanMinimum(userId string) (*[]dto.FoodDto, error)
}

type PurchaseService struct {
	pr repositories.PurchaseRepositoryInterface
}

func NewPurchaseService(pr repositories.PurchaseRepositoryInterface) *PurchaseService {
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
	for _, food := range foods {
		purchase.TotalCost += food.UnitPrice * (float64)(food.MinimumQuantity-food.CurrentQuantity)
		purchase.Foods = append(purchase.Foods, models.PurchaseQuantity{
			FoodCode: utils.GetStringIDFromObjectID(food.Code),
			Quantity: food.MinimumQuantity - food.CurrentQuantity,
		})
	}
	res, err := ps.pr.Create(purchase)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ps *PurchaseService) GetFoodWithQuantityLessThanMinimum(userId string) (*[]dto.FoodDto, error) {
	foods, err := ps.pr.GetFoodWithQuantityLessThanMinimum(userId)
	if err != nil {
		return nil, err
	}
	var foodDtos []dto.FoodDto
	for _, food := range foods {
		foodDtos = append(foodDtos, dto.FoodDto{
			Code:            utils.GetStringIDFromObjectID(food.Code),
			Name:            food.Name,
			UnitPrice:       food.UnitPrice,
			CurrentQuantity: food.CurrentQuantity,
			MinimumQuantity: food.MinimumQuantity,
		})
	}

	return &foodDtos, nil
}


