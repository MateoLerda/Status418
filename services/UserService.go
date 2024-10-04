package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
)

type UserServiceInterface interface {
	GetAll() ([]dto.UserDTO, error)
	GetById(id int) (dto.UserDTO, error)
	Create(dto.UserDTO) error
	Update(dto.UserDTO) error
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

func (us *UserService) GetAll() (*[]dto.UserDTO, error) {
	users, err := us.ur.GetAll()
	if err != nil {
		return nil, err
	}
	usersDTO := ChangeFromModelToDto(users)
	return usersDTO, nil

}
func ChangeFromModelToDto(users *[]models.User) *[]dto.UserDTO {
	var usersDTO []dto.UserDTO
	for _, user := range *users {
		userDTO := dto.NewUserDTO(user.Name, user.LastName, user.Email, user.Password)
		usersDTO = append(usersDTO, *userDTO)
	}
	return &usersDTO
}
