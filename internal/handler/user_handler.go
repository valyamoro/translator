package handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/valyamoro/internal/domain"

	"fmt"
)

type Users interface {
	Create(user domain.User) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
}

type UserHandler struct {
	UsersService Users
	Validator *validator.Validate
}

func NewUserHandler(users Users) *UserHandler {
	v := validator.New()
	
	uh := &UserHandler{
		UsersService: users,
		Validator: v,
	}
	
	v.RegisterValidation("exists_user", uh.userExists)
	v.RegisterValidation("passwd", validatePassword)

	return uh
}

func (uh *UserHandler) userExists(fl validator.FieldLevel) bool {
    username := fl.Field().String()
    _, err := uh.UsersService.GetByUsername(username)

    if err == nil {
        return false 
    }

    return false 
}

func (uh *UserHandler) InitRoutes(router *gin.Engine) {
	router.POST("/users", uh.CreateUser)
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#~$*(),.?"{}|<>]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial
}


func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.Validator.Struct(user); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []string 
			
			for _, validationErr := range validationErrors {
				switch validationErr.Tag() {
				case "required":
					errors = append(errors, fmt.Sprintf("%s is required", validationErr.Field()))
				case "min":
					errors = append(errors, fmt.Sprintf("%s must be at least %s characters long", validationErr.Field(), validationErr.Param()))
				case "passwd":
					errors = append(errors, "Password must be at least 8 characters long and contain a mix of upper, lower, digits, and special characters")
				default:
					errors = append(errors, fmt.Sprintf("Invalid field %s", validationErr.Field()))
				}
			}

			c.JSON(http.StatusBadRequest, gin.H{"validation_error": errors})
			return 
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error occured"})
		return 
	}

	_, err := uh.UsersService.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
