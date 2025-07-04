package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadProperty(propertyName string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	propertyValue := os.Getenv(propertyName)
	return propertyValue
}
