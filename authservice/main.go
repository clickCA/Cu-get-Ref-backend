package main

import (
	"authservice/config"
	"authservice/controllers"
	_ "authservice/docs"
	"authservice/models"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func main() {
	// Load .env file

	db, err := config.ConnectDB()
	// Migrate the schema
	db.AutoMigrate(&models.User{})
	// Initialize a Zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync() // Flushes buffer before exit

	// Create a new Gin router
	r := gin.Default()

	// Create controllers for the signup and signin handlers
	signupController := controllers.NewRegisterController(logger)
	signinController := controllers.NewLoginController(logger)

	// Define the signup and signin routes
	r.POST("/register", gin.WrapF(signupController.RegisterHandler))
	r.POST("/login", gin.WrapF(signinController.LoginHandler))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))
	// Start the server

	logger.Info("Server started on :8080")
	r.Run(":8080")
}
