package environment

import (
	"os"
	"strconv"
)

func Get(key string) string {
	return os.Getenv(key)
}

func GetOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable '" + key + "' is required")
	}
	return value
}

func GetIntOrDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intVal
}

func GetIntOrPanic(key string) int {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable '" + key + "' is required")
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		panic("Environment variable '" + key + "' is required to be an integer")
	}
	return intVal
}
