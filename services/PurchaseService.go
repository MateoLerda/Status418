package services

import (
	"Status418/dto"
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
	pr repositories.PurchaseRepositoryInterface
}

func NewPurchaseService(pr repositories.PurchaseRepositoryInterface) *PurchaseService {
	return &PurchaseService{
		pr: pr,
	}
}

func calculatePurchaseAllFoods(foods []models.Food) models.Purchase {
	var purchase models.Purchase
	for _, food := range foods {
		purchase.TotalCost += food.UnitPrice * (float64)(food.MinimumQuantity-food.CurrentQuantity)
		purchase.Foods = append(purchase.Foods, models.PurchaseQuantity{
			FoodCode: utils.GetStringIDFromObjectID(food.Code),
			Quantity: food.MinimumQuantity - food.CurrentQuantity,
		})
	}
	return purchase
}

func (ps *PurchaseService) Create(userId string, purchaseDto dto.PurchaseDto) (*mongo.InsertOneResult, error) {
	DB := repositories.NewMongoDB()
	foodRepository := repositories.NewFoodRepository(DB)
	var foods []models.Food
	var err error
	var purchase models.Purchase

	if len(purchaseDto.Foods) == 0 {
		foods, err = foodRepository.GetAll(userId, true)
		if err != nil {
			return nil, err
		}
		purchase = calculatePurchaseAllFoods(foods)
	} else {
		var food models.Food
		for _, purchaseQuantity := range purchaseDto.Foods {
			food, err = foodRepository.GetByCode(purchaseQuantity.FoodCode, userId)
			if err != nil {
				return nil, err
			}
			purchase.TotalCost += food.UnitPrice * float64(purchaseQuantity.Quantity)
			purchase.Foods = append(purchase.Foods, models.PurchaseQuantity{FoodCode: utils.GetStringIDFromObjectID(food.Code), Quantity: purchaseQuantity.Quantity})
		}
	}
	
	purchase.PurchaseDate = time.Now()
	purchase.UserId= utils.GetObjectIDFromStringID(userId)
	res, err := ps.pr.Create(purchase)
	if err != nil {
		return nil, err
	}
	return res, nil
}
