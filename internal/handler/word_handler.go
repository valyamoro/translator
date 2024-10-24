package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/valyamoro/internal/constants"
	"github.com/valyamoro/internal/domain"
	"regexp"
)

type Words interface {
	Create(word domain.Word) (domain.Word, error)
	GetByID(id int64) (domain.Word, error)
	GetAll() ([]domain.Word, error)
	Delete(id int64) (domain.Word, error)
	Update(id int64, word domain.Word) (domain.Word, error)
}

type WordHandler struct {
	BaseHandler
	wordsService Words
	dictionaryService Dictionaries
	Validator    *validator.Validate
}

func NewWordHandler(words Words, dictionaries Dictionaries) *WordHandler {
	v := validator.New()

	wh := &WordHandler{
		wordsService: words,
		Validator:    v,
		dictionaryService: dictionaries,
	}

	v.RegisterValidation("language", isValidLanguage)
	v.RegisterValidation("correct", isCorrectWord)
	v.RegisterValidation("dictionary_exists", wh.dictionaryExists)
	v.RegisterStructValidation(isIncomingCodesIdentical, domain.Word{})
	
	return wh
}

func isValidLanguage(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	switch constants.Language(value) {
	case constants.Russian, constants.English:
		return true 
	default:
		return false
	}
}

func isCorrectWord(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	re := regexp.MustCompile(`[!@#$^&*()]`)

	return !re.MatchString(value)
}

func (wh *WordHandler) dictionaryExists(fl validator.FieldLevel) bool {
	value := fl.Field().Int()

	_, err := wh.dictionaryService.GetByID(value)

	if err == nil {
		return true 
	} else {
		return false
	}
}

func isIncomingCodesIdentical(sl validator.StructLevel) {
	word := sl.Current().Interface().(domain.Word)

	if word.WordLanguageCode == word.TranslationWordLanguageCode {
		sl.ReportError(
			word.WordLanguageCode,      
			"WordLanguageCode",      
			"WordLanguageCode",          
			"identicalLanguageCodes",
			"",
		)
	}
}

func (wh *WordHandler) InitRoutes(router *gin.Engine) {
	router.GET("/words", wh.GetAllWords)
	router.GET("/words/:id", wh.GetWordByID)
	router.POST("/words", wh.CreateWord)
	router.PUT("/words/:id", wh.UpdateWord)
	router.DELETE("/words/:id", wh.DeleteWord)
}

func (wh *WordHandler) CreateWord(c *gin.Context) {
	var word domain.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := wh.Validator.Struct(word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := wh.wordsService.Create(word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, word)
}

func (wh *WordHandler) GetWordByID(c *gin.Context) {
	id, err := wh.BaseHandler.getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	word, err := wh.wordsService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, word)
}

func (wh *WordHandler) GetAllWords(c *gin.Context) {
	words, err := wh.wordsService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, words)
}

func (wh *WordHandler) DeleteWord(c *gin.Context) {
	id, err := wh.BaseHandler.getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = wh.wordsService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (wh *WordHandler) UpdateWord(c *gin.Context) {
	id, err := wh.BaseHandler.getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var word domain.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := wh.Validator.Struct(word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	word.ID = id 

	_, err = wh.wordsService.Update(id, word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, word)
}
