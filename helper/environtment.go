package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func Env() {
	err := godotenv.Load(".env")
	PanicIfError(err)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
