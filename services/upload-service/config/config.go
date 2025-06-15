package config

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found. Using system env instead")
	}
}