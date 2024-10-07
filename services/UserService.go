package services

import (
	"Status418/dto"
	//"Status418/models"
	"Status418/repositories"
	"time"
)
//REVISAR TODO ESTE ARCHIVO. HAY COSAS QUE NO ESTAN CORRECTAS EN ALGUNOS MÉTODOS YA QUE NO SE ESTÁN USANDO LOS PARÁMETROS TAL CUAL
//APARECIAN EN LA INTERFAZ

type UserServiceInterface interface {
	GetAll() ([]dto.UserDto, error)
	GetById(id string) (dto.UserDto, error)
	Create(dto.UserDto) error
	Update(dto.UserDto) error
	Delete(id string) error
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
	var usersDTO []dto.UserDto
	users, err := us.ur.GetAll()
	if err != nil {
		return nil, err
	}
	for _, user := range *users {
		userDTO := dto.NewUserDto(user)
		usersDTO = append(usersDTO, *userDTO)
	}
	return &usersDTO, nil
}

func (us *UserService) GetById(id string) (*dto.UserDto, error) {
	user, err := us.ur.GetById(id)
	if err != nil {
		return nil, err
	}
	userDTO := dto.NewUserDto(*user)
	return userDTO, nil
}

func (us *UserService) Create(userDTO dto.UserDto) error {
	user := userDTO.GetModel()
	user.CreationDate = time.Now().String()
	_, err := us.ur.Create(&user)
	if err != nil {
		return err
	}
	return nil
}
func (us *UserService) Update(userDTO dto.UserDto) error {
	user := userDTO.GetModel()
	_, err := us.ur.Update(&user)
	if err != nil {
		return err
	}
	return nil
}
func (us *UserService) Delete(id string) error {
	_, err := us.ur.Delete(id)
	if err != nil {
		return err
	}
	return nil
}


