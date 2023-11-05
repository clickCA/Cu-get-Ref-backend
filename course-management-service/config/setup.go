package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// Get database credentials from .env file
	dsn := EnvMySqlURI()

	if dsn == "" {
		log.Fatal("DB_HOST not found in .env file")
	}

	// Connect to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}

var DB = ConnectDB()
