package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

func getString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	val, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return val
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	val, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return val
}

type DatabaseConfig struct {
	url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: getString("DATABASE_URL", ""),
	}
}
