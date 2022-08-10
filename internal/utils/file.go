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

func IsDirExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
