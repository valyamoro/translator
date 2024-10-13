package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/valyamoro/internal/handler"
	"github.com/valyamoro/internal/repository"
	"github.com/valyamoro/internal/service"
	"github.com/valyamoro/pkg/database"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte("secret-key")

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

	userHandler.InitRoutes(router)
	dictionaryHandler.InitRoutes(router)
	wordHandler.InitRoutes(router)

	router.Run(":8080")
}

func GenerateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil 
	} else {
		return nil, err 
	}
}
