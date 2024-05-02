package initializers

import (
	"log"

	"github.com/lpernett/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file ❌📝 : %s", err)
	}
}
