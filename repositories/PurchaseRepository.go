package repositories

import "Status418/models"

type PurchaseRepository interface {
	GetFoodWithQuantityLessThanMinimum() ([]models.Food,error)
	Create(models.Purchase)error
}
