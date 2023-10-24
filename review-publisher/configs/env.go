package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvAmqpURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: AMQPURI")
	}

	return os.Getenv("AMQPURI")
}
