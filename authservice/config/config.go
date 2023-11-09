package config

import (
	"authservice/models"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	JWT_KEY    string `mapstructure:"JWT_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env.local")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func ConnectDB() (*gorm.DB, error) {
	config, err := LoadConfig("./config")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.User{})

	return db, nil
}

func GetJWTSecret() ([]byte, string) {
	config, err := LoadConfig("./config")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}
	return []byte(config.JWT_SECRET), config.JWT_KEY
}
