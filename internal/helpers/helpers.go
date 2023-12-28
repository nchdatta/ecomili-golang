package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GetIntFromEnv(key string) int {
	valStr := os.Getenv(key)
	valInt := 0
	if valStr != "" {
		convertedVal, err := strconv.Atoi(valStr)
		if err == nil {
			valInt = convertedVal
		}
	}
	return valInt
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
