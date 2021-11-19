package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var PORT int

func LoadEnv(key string) string{
	err := godotenv.Load(".env")
	if err != nil {
		return "error fetching environment variable"
	}
	return os.Getenv(key)
}

func InitPort() {
	PORT, _ = strconv.Atoi(LoadEnv("PORT"))
}