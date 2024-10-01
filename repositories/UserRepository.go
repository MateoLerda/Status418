package repositories

import "Status418/models"

type UserRepository interface {
	GetAll()([]models.User, error)
	GetById(id int)(models.User, error)
	Create(models.User) error
	Update(models.User) error
	Delete(id int) error
}