package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/valyamoro/internal/handler"
	"github.com/valyamoro/internal/repository"
	"github.com/valyamoro/internal/service"
	"github.com/valyamoro/pkg/database"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инициализации конфигурации: %v", err)
		return
	}

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

	router.Run(":" + viper.GetString("PORT"))
}

func initConfig() error {
    viper.SetConfigFile(".env")
    viper.AddConfigPath("../") 
    viper.AddConfigPath(".")
	viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Не удалось найти конфиг: %v", err)
        return err
    }

    return nil
}

func initDB() (*sql.DB, error) {
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	username := viper.GetString("DB_USERNAME")
	dbName := viper.GetString("DB_NAME")
	sslMode := viper.GetString("DB_SSLMODE")
	password := viper.GetString("DB_PASSWORD")

	return database.NewPostgresConnection(database.ConnectionParams{
		Host:     host,
		Port:     port,
		Username: username,
		DBName:   dbName,
		SSLMode:  sslMode,
		Password: password,
	})
}
