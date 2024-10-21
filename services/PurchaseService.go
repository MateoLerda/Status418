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
	Create(userCode string, purchaseDto dto.PurchaseDto) (*mongo.InsertOneResult, error)
}

type PurchaseService struct {
	purchaseRepository repositories.PurchaseRepositoryInterface
}

func NewPurchaseService(purchaseRepository repositories.PurchaseRepositoryInterface) *PurchaseService {
	return &PurchaseService{
		purchaseRepository: purchaseRepository,
	}
}

func calculatePurchaseAllFoods(foods []models.Food) models.Purchase {
	var purchase models.Purchase
	for _, food := range foods {
		purchase.TotalCost += food.UnitPrice * (float64)(food.MinimumQuantity-food.CurrentQuantity)
		purchase.Foods = append(purchase.Foods, models.FoodQuantity{
			FoodCode: utils.GetStringIDFromObjectID(food.Code),
			Quantity: food.MinimumQuantity - food.CurrentQuantity,
		})
	}
	return purchase
}

func (purchaseService *PurchaseService) Create(userCode string, purchaseDto dto.PurchaseDto) (*mongo.InsertOneResult, error) {
	DB := repositories.NewMongoDB()
	foodRepository := repositories.NewFoodRepository(DB)
	var foods []models.Food
	var err error
	var purchase models.Purchase

	if len(purchaseDto.Foods) != 0 {
		var food models.Food
		for _, foodQuantity := range purchaseDto.Foods {
			food, err = foodRepository.GetByCode(foodQuantity.FoodCode, userCode)
			if err != nil {
				return nil, err
			}
			purchase.TotalCost += food.UnitPrice * float64(foodQuantity.Quantity)
			purchase.Foods = append(purchase.Foods, models.FoodQuantity{FoodCode: foodQuantity.FoodCode, Quantity: foodQuantity.Quantity})
		}
	} else {
		foods, err = foodRepository.GetAll(userCode, true)
		if err != nil {
			return nil, err
		}
		purchase = calculatePurchaseAllFoods(foods)
	}

	purchase.PurchaseDate = time.Now()
	purchase.UserCode = userCode
	res, err := purchaseService.purchaseRepository.Create(purchase)
	if err != nil {
		return nil, err
	}
	return res, nil
}
