package main

import (
	"fmt"
	"log"
	"os"
	_ "profileservice/docs"
	"profileservice/internal/handlers"
	"profileservice/internal/models"
	"profileservice/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
	fmt.Print(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Student{}, &models.Professor{}, &models.Course{}, &models.Review{})

	// Initialize and use the auth service
	profileService := services.New(db)

	// Initialize the router
	router := gin.Default()
	router.Use(CORSMiddleware())

	profileHandler := handlers.NewProfileHandler(profileService)

	router.POST("/profiles", profileHandler.Create)
	router.GET("/profiles/:id", profileHandler.Read)
	router.PUT("/profiles/:id", profileHandler.Update)
	router.DELETE("/profiles/:id", profileHandler.Delete)

	// swagger definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8081") // starts the gin server
}
