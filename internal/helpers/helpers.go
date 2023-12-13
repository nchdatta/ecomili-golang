package helpers

import (
	"os"
	"strconv"
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
