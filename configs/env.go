package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file fro mongo")
    }
  
    return os.Getenv("MONGOURI")
}

func GetPort() string {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file for port")
    }
  
    return os.Getenv("PORT")
}