package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Config contains our config struct.
type Config struct {
	// KeySize is the size of the RSA key in bytes. 1024, 2048 or 4096
	KeySize int
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment or return a default value
func getEnvRequired(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatal("ERROR: required option ", key)
	return ""
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
