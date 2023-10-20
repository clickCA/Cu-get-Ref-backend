package main

import (
	"authservice/controllers"
	_ "authservice/docs"
	"authservice/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

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

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Create controllers for the signup and signin handlers
	signupController := controllers.NewSignupController(logger)
	signinController := controllers.NewSigninController(logger)

	// Define the signup and signin routes
	r.HandleFunc("/signup", signupController.SignupHandler).Methods("POST")
	r.HandleFunc("/signin", signinController.SigninHandler).Methods("POST")
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs"))))

	// Start the server
	http.Handle("/", r)
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Info("Server started on :" + port)
	http.ListenAndServe(":"+port, nil)
}
func ServeSwagger(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./docs/index.html")
}
