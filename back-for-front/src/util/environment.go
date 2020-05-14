package util

import (
	"os"
	"strings"
)

// Address represents address of each environment
var Address = map[string]string{
	"dev":  "localhost:8081",
	"prod": ":8081",
}

// Getenv find environment variable on SO and return default value if not found
func Getenv(variable, defaultValue string) string {
	value := os.Getenv(variable)

	if strings.EqualFold(value, "") {
		return defaultValue
	}
	return value
}
