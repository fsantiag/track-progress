package util

import (
	"os"
	"strings"
)

// Getenv find environment variable on SO and return default value if not found
func Getenv(variable, defaultValue string) string {
	value := os.Getenv(variable)

	if strings.EqualFold(value, "") {
		return defaultValue
	}
	return value
}
