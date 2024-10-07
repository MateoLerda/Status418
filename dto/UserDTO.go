package dto

import (
	"Status418/models"
	"Status418/utils"
)

type UserDto struct {
	UserId   string
	Name     string
	LastName string
	Email    string
	Password string
}

func NewUserDto(model models.User) *UserDto {
	return &UserDto{
		UserId:   utils.GetStringIDFromObjectID(model.UserId),
		Name:     model.Name,
		LastName: model.LastName,
		Email:    model.Email,
		Password: model.Password,
	}
}

func (dto UserDto) GetModel() models.User {
	return models.User{
		UserId:   utils.GetObjectIDFromStringID(dto.UserId),
		Name:     dto.Name,
		LastName: dto.LastName,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
