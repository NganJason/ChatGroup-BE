package utils

import (
	"os"
)

const (
	envFilePath = ".env"
)

func GetDotEnvVariable(key string) (string, error) {
	return os.Getenv(key), nil
}
