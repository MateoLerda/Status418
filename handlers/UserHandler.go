package handlers

import "Status418/services"

type UserHandler struct {
	us services.UserServiceInterface
}

func NewUserHandler(us services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

//IMPLEMENTAR LOS MÉTODOS DE LA INTERFACE UserServiceInterfaces
