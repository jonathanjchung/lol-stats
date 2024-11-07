package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Couldn't load .env file")
	}

	return os.Getenv(key)
}
