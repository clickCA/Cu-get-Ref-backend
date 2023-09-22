package main

import (
	_ "REST/docs"
	"REST/internal/auth"
	"REST/internal/handlers"
	"REST/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	// Formulate the connection string using the loaded environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Student{}, &models.Professor{}, &models.Course{}, &models.Review{})

	// Initialize and use the auth service
	authService := auth.New(db)

	// Initialize the router
	router := gin.Default()

	authHandler := handlers.NewAuthHandler(authService)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	// swagger definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080") // starts the gin server
}
