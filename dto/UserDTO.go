package dto

type UserDto struct {
	Name     string
	LastName string
	Email    string
	Password string
}

func NewUserDTO(Name string, LastName string, Email string, Password string) *UserDto {
	return &UserDto{
		Name:     Name,
		LastName: LastName,
		Email:    Email,
		Password: Password,
	}
}
