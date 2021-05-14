// package config contains the configuration options and structs.
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
	keySize int
}

func NewConfig() *Config {
	return &Config{
		keySize: getEnvAsInt("RSAKEYBYTESIZE", 2048),
	}
}

// KeySize returns the KeySize
func (c *Config) KeySize() int {
	return c.keySize
}

// getEnv is a simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// getEnvRequired is a simple helper function to read an environment or return a default value
func getEnvRequired(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatal("ERROR: required option ", key)
	return ""
}

// getEnvAsInt is a simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// getEnvAsSlice is a helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
