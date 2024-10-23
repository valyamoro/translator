package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/valyamoro/internal/domain"
	"net/http"
)

type Dictionaries interface {
	Create(dictionary domain.Dictionary) (domain.Dictionary, error)
	GetByID(id int64) (domain.Dictionary, error)
	GetAll() ([]domain.Dictionary, error)
	Delete(id int64) (domain.Dictionary, error)
	Update(id int64, inp domain.UpdateDictionaryInput) (domain.Dictionary, error)
}

type DictionaryHandler struct {
	BaseHandler
	Validator *validator.Validate
	dictionariesService Dictionaries
	usersService Users
}

func NewDictionaryHandler(dictionaries Dictionaries, users Users) *DictionaryHandler {
	v := validator.New()

	dh := &DictionaryHandler{
		dictionariesService: dictionaries,
		usersService: users,
		Validator: v,
	}

	v.RegisterValidation("user_exists", dh.userExists)

	return dh 
}

func (dh *DictionaryHandler) InitRoutes(router *gin.Engine) {
	router.GET("/dictionaries", dh.GetAllDictionaries)
	router.GET("/dictionaries/:id", dh.GetDictionaryByID)
	router.POST("/dictionaries", dh.CreateDictionary)
	router.PUT("/dictionaries/:id", dh.UpdateDictionary)
	router.DELETE("/dictionaries/:id", dh.DeleteDictionary)
}

func (dh *DictionaryHandler) userExists(fl validator.FieldLevel) bool {
	value := fl.Field().Int()

	_, err := dh.usersService.GetByID(int(value))
	if err == nil {
		return true 
	} else {
		return false
	}
}

func (dh *DictionaryHandler) CreateDictionary(c *gin.Context) {
	var dictionary domain.Dictionary
	if err := c.ShouldBindJSON(&dictionary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dh.Validator.Struct(dictionary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	_, err := dh.dictionariesService.Create(dictionary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dictionary)
}

func (dh *DictionaryHandler) GetDictionaryByID(c *gin.Context) {
	id, err := dh.BaseHandler.getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dictionary, err := dh.dictionariesService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dictionary)
}

func (dh *DictionaryHandler) GetAllDictionaries(c *gin.Context) {
	dictionaries, err := dh.dictionariesService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dictionaries)
}

func (dh *DictionaryHandler) DeleteDictionary(c *gin.Context) {
	id, err := dh.BaseHandler.getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = dh.dictionariesService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (dh *DictionaryHandler) UpdateDictionary(c *gin.Context) {
	id, err := dh.BaseHandler.getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inp domain.UpdateDictionaryInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = dh.dictionariesService.Update(id, inp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
