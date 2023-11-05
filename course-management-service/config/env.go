package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMySqlURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: DB_HOST")
	}

	return os.Getenv("DB_HOST")
}

func EnvServerPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: SERVER_PORT")
	}

	return os.Getenv("SERVER_PORT")
}
