package main

import (
	"context"
	"github.com/NorthDice/ReflectDiary/internal/adapter/controllers/http/handlers"
	"github.com/NorthDice/ReflectDiary/internal/infrastructure/repository/postgres"
	"github.com/NorthDice/ReflectDiary/internal/infrastructure/services"
	"github.com/NorthDice/ReflectDiary/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// @title Reflect Diary API
// @version 1.0
// @description API Server for ReflectDiary application

// @host localhost: 8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	router := gin.Default()
	ctx := context.Background()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	}

	db, err := postgres.NewPostgresRepository(ctx, cfg)
	if err != nil {
		panic(err)
	}

	userRepository := postgres.NewUserPostgresRepository(db)

	passwordService := services.NewPasswordService()
	authService := services.NewAuthService(userRepository)

	registerUseCase := user.NewRegisterUseCase(userRepository, passwordService, authService)
	loginUseCase := user.NewLoginUseCase(userRepository, passwordService, authService)
	userHandler := handlers.NewUserHandler(registerUseCase, loginUseCase)

	handlers.InitRoutes(router, userHandler)
}
