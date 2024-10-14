package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/valyamoro/internal/handler"
	"github.com/valyamoro/internal/repository"
	"github.com/valyamoro/internal/service"
	"github.com/valyamoro/pkg/database"

	"strconv"
)


func main() {
	initConfig()

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
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

	userHandler.InitRoutes(router)
	dictionaryHandler.InitRoutes(router)
	wordHandler.InitRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
}

func initDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	username := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")
	password := os.Getenv("DB_PASSWORD")

	return database.NewPostgresConnection(database.ConnectionParams{
		Host:     host,
		Port:     port,
		Username: username,
		DBName:   dbName,
		SSLMode:  sslMode,
		Password: password,
	})
}
