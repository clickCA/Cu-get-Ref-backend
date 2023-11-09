package config

import (
	"course-management-service/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// Get database credentials from .env file
	dsn := EnvMySqlURI()

	// Connect to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = db.AutoMigrate(&models.Course{})
	if err != nil {
		log.Fatal("Error migrating database")
	}

	return db
}

var DB = ConnectDB()
