pakage config


import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func config() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("DB_HOST")
	fmt.Println(secretKey)
}