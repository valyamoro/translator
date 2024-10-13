package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/valyamoro/pkg/jwt"
	"github.com/valyamoro/internal/domain"
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
	router.POST("/create", uh.CreateUser)
	router.POST("/login", uh.Login)
}

func (h *UserHandler) Login(c *gin.Context) {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := jwt.GenerateJWT(loginData.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
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
