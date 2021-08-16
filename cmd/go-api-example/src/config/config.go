// Config contains functions that are used to define app config.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// List of used env variables.
var (
	// Default Port
	PORT = getEnv("PORT", "3000")
	// JWT token secret
	JWT_SECRET = getEnv("JWT_SECRET", "")
)

// Initialize initialize middlewares.
func Initialize(a fiber.App) {
	a.Use(logger.New())
}

// getEnv reads env file and get variable by given name.
func getEnv(n, f string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while loading env file \n%v", err)
	}

	if value := os.Getenv(n); value != "" {
		return value
	}

	if f != "" {
		return f
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, n))
}
