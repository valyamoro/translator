package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyamoro/internal/handler"
	"github.com/valyamoro/internal/repository"
	"github.com/valyamoro/internal/service"
	"github.com/valyamoro/pkg/database"
	"os"
	"strconv"
)

func main() {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	username := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")
	password := os.Getenv("DB_PASSWORD")

	db, err := database.NewPostgresConnection(database.ConnectionParams{
		Host:     host,
		Port:     port,
		Username: username,
		DBName:   dbName,
		SSLMode:  sslMode,
		Password: password,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	userRepo := repository.NewUsersRepository(db)
	dictionaryRepo := repository.NewDictionaryRepository(db)
	wordRepo := repository.NewWordsRepository(db)

	userService := service.NewUsersService(userRepo)
	dictionaryService := service.NewDictionariesService(dictionaryRepo)
	wordService := service.NewWordsService(wordRepo)

	userHandler := handler.NewUserHandler(userService)
	dictionaryHandler := handler.NewDictionaryHandler(dictionaryService)
	wordHandler := handler.NewWordHandler(wordService)

	router := gin.Default()

	router.POST("/users", userHandler.CreateUser)

	router.GET("/dictionaries", dictionaryHandler.GetAllDictionaries)
	router.GET("/dictionaries/:id", dictionaryHandler.GetDictionaryByID)
	router.POST("/dictionaries", dictionaryHandler.CreateDictionary)
	router.PUT("/dictionaries/:id", dictionaryHandler.UpdateDictionary)
	router.DELETE("/dictionaries/:id", dictionaryHandler.DeleteDictionary)

	router.GET("/words", wordHandler.GetAllWords)
	router.GET("/words/:id", wordHandler.GetWordByID)
	router.POST("/words", wordHandler.CreateWord)
	router.PUT("/words/:id", wordHandler.UpdateWord)
	router.DELETE("/words/:id", wordHandler.DeleteWord)

	router.Run(":8080")
}
