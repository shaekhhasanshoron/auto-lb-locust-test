package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ServerPort refers to server port.
var ServerPort string


// InitEnvironmentVariables initializes environment variables
func InitEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err.Error(), ", reading from env")

	}
	ServerPort = os.Getenv("SERVER_PORT")
}
