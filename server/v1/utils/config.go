package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string `json:"port"`
	DbUrl string `json:"dbUrl"`
}

func GetConfig() Config {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	return Config{
		Port:  port,
		DbUrl: dbUrl,
	}
}
