package main

import (
	"authservice/controllers"
	_ "authservice/docs"
	"authservice/middleware"
	"authservice/models"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
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
	// Initialize a Zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync() // Flushes buffer before exit

	tokenMiddleware := middleware.NewTokenMiddleware(logger)
	// Create a new Gin router
	r := gin.Default()
	r.Use(tokenMiddleware.TokenValidationMiddleware())

	// Create controllers for the signup and signin handlers
	signupController := controllers.NewSignupController(logger)
	signinController := controllers.NewSigninController(logger)

	// Define the signup and signin routes
	r.POST("/signup", gin.WrapF(signupController.SignupHandler))
	r.POST("/signin", gin.WrapF(signinController.SigninHandler))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))
	// Start the server
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Info("Server started on :" + port)
	r.Run(":" + port)
}
