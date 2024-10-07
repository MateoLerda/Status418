package handlers

import (
	"Status418/dto"
	"Status418/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	us services.UserServiceInterface
}

func NewUserHandler(us services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (uh *UserHandler) GetAll(c *gin.Context) {
	users, err := uh.us.GetAll()

	if err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users from database",
		})
		return
	}

	if err.Error() == "notfound" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found any users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetById(c *gin.Context) {
	userId := c.Param("userId")
	user, err := uh.us.GetById(userId)

	if err.Error() == "internal" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user with ID: " + userId,
		})
		return
	}

	if err.Error() == "notfound" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found any user with ID: " + userId,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) Create(c *gin.Context) {
	var user dto.UserDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uh.us.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user. " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func Update() {

}

func Delete() {

}
