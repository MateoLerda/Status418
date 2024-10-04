package dto

type UserDTO struct {

	Name string 
	LastName string 
	Email string 
	Password string
}

func NewUserDTO(Name string, LastName string, Email string, Password string) *UserDTO {
	return &UserDTO{
		Name: Name,
		LastName: LastName,
		Email: Email,
		Password: Password,
	}
}