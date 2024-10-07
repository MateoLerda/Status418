package dto
 
import "Status418/models"

type UserDto struct {
	UserID   string
	Name     string
	LastName string
	Email    string
	Password string
}

func NewUserDto(model models.User) *UserDto {
	return &UserDto{
		Name:     Name,
		LastName: LastName,
		Email:    Email,
		Password: Password,
	}
}

func (dto UserDto) GetModel() models.User {
	return models.User{
		Name:     dto.Name,
		LastName: dto.LastName,
		Email:    dto.Email,
		Password: dto.Password,
		CreationDate: 
		UpdateDate:
	}
	
}
