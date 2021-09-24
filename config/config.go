package config

import (
	"os"
	"strconv"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	num, _ := strconv.Atoi(os.Getenv(key))
	return num
}

func GetBool(key string) bool {
	val := os.Getenv(key)
	return val == "1" || val == "true" || val == "TRUE"
}
