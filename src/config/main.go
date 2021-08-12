// package to define config
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// List of used env variables
var (
	// Default Port
	PORT = getEnv("PORT", "3000")
	// JWT token secret
	JWT_SECRET = getEnv("JWT_SECRET", "")
)

// Read env file and get variable by given name
func getEnv(name string, fallback string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while loading env file \n%v", err)
	}

	if value := os.Getenv(name); value != "" {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
