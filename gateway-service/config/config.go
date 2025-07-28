package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	//load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
