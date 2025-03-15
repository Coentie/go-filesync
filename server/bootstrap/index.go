package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
)

func Boostrap() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	return nil
}
