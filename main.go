// Package main implements web server.
// Here is the server starts
package main

import (
	"github.com/2rueSid/go-api-example/src/config"
	"github.com/2rueSid/go-api-example/src/config/database"
	"github.com/2rueSid/go-api-example/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Get fiber application instance
	app := fiber.New()

	// Initialize middlewares
	config.Initialize(*app)

	// Initialize routes
	routes.Initialize(app)

	// Run application on given PORT
	// Default PORT equals to 3000
	app.Listen(config.PORT)

	// Close database connection
	defer database.Disconnect()
}
