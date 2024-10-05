package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
)

type UserServiceInterface interface {
	GetAll() ([]dto.UserDto, error)
	GetById(id int) (dto.UserDto, error)
	Create(dto.UserDto) error
	Update(dto.UserDto) error
	Delete(id int) error
}

type UserService struct {
	ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) GetAll() (*[]dto.UserDto, error) {
	users, err := us.ur.GetAll()
	if err != nil {
		return nil, err
	}
	usersDTO := ChangeFromModelToDto(users)
	return usersDTO, nil

}

func ChangeFromModelToDto(users *[]models.User) *[]dto.UserDto {
	var usersDto []dto.UserDto
	for _, user := range *users {
		userDto := dto.NewUserDTO(user.Name, user.LastName, user.Email, user.Password)
		usersDto = append(usersDto, *userDto)
	}
	return &usersDto
}
