package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	if os.Args[len(os.Args)-1] == "dev" {
		err := godotenv.Load("dev.env")
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	fmt.Println("Environment variables Loaded")
}

func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
