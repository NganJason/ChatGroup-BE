package utils

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	envFilePath = ".env"
)

func GetDotEnvVariable(key string) (string, error) {
	godotenv.Load(envFilePath)

	return os.Getenv(key), nil
}
