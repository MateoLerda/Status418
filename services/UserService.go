package services

import "Status418/dto"

type UserServiceInterface interface {
	GetAll() ([]dto.UserDTO, error)
	GetById(id int) (dto.UserDTO, error)
	Create(dto.UserDTO) error
	Update(dto.UserDTO) error
	Delete(id int) error
}