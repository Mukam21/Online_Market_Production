package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/Mukam21/Online_Market_Production/docs"
	"github.com/Mukam21/Online_Market_Production/internal/database"
	"github.com/Mukam21/Online_Market_Production/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Online Market API
// @version 1.0
// @description API для онлайн магазина
// @host localhost:8080
// @BasePath /api

func main() {
	// Загрузка .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file")
	}

	// Подключение к БД
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := database.NewDatabase(dsn)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Инициализация роутера
	router := gin.Default()

	// Регистрация маршрутов
	routes.SetupRoutes(router, db.DB)

	// Запуск сервера
	router.Run(":8080")
}
