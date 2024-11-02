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
	Create(userCode string, newPurchase dto.PurchaseDto) (*mongo.InsertOneResult, error)
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
			FoodCode: food.Code,
			Name: food.Name,
			Quantity: food.MinimumQuantity - food.CurrentQuantity,
		})
	}
	return purchase
}

func (purchaseService *PurchaseService) Create(userCode string, newPurchase dto.PurchaseDto) (*mongo.InsertOneResult, error) {
	DB := repositories.NewMongoDB()
	foodRepository := repositories.NewFoodRepository(DB)
	var foods []models.Food
	var err error
	var purchase models.Purchase

	if len(newPurchase.Foods) != 0 {
		var food models.Food
		for _, foodQuantity := range newPurchase.Foods {
			foodObjectId := utils.GetObjectIDFromStringID(foodQuantity.FoodCode)
			food, err = foodRepository.GetByCode(foodObjectId, userCode)
			if err != nil {
				return nil, err
			}
			purchase.TotalCost += food.UnitPrice * float64(foodQuantity.Quantity)
			purchase.Foods = append(purchase.Foods, models.FoodQuantity{FoodCode: foodObjectId, Name: food.Name, Quantity: foodQuantity.Quantity})
		}
	} else {
		foods, err = foodRepository.GetAll(userCode, true)
		if err != nil {
			return nil, err
		}
		purchase = calculatePurchaseAllFoods(foods)
	}

	purchase.PurchaseDate = time.Now().String()
	purchase.UserCode = userCode
	create, err := purchaseService.purchaseRepository.Create(purchase)
	if err != nil {
		return nil, err
	}
	for _, food := range purchase.Foods {
		var updatedFood models.Food
		updatedFood.Code = food.FoodCode
		updatedFood.CurrentQuantity = food.Quantity
		_,err := foodRepository.Update(updatedFood, false)
		if err!= nil {
			return nil, err
		}
	}
	return create, nil
}
//ahora nop porque ahora si ante no ahora