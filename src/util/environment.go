package util

import "os"

// Getenv find environment variable on SO and return default value if not found
func Getenv(variable, defaultValue string) string {
	value := os.Getenv(variable)
	if value == "" {
		return defaultValue
	}
	return value
}
