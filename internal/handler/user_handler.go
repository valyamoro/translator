package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/valyamoro/internal/domain"
	"net/http"
)

type Users interface {
	Create(user domain.User) (domain.User, error)
}

type UserHandler struct {
	UsersService Users
}

func NewUserHandler(users Users) *UserHandler {
	return &UserHandler{
		UsersService: users,
	}
}

func (uh *UserHandler) InitRoutes(router *gin.Engine) {
	router.POST("/users", uh.CreateUser)
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := uh.UsersService.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
