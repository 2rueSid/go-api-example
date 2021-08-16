// Main implements web server.
package main

import (
	"log"

	"github.com/2rueSid/go-api/cmd/go-api-example/src/config"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/config/database"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Get fiber application instance.
	a := fiber.New()

	// Initialize middlewares.
	config.Initialize(*a)

	// Initialize routes.
	routes.Initialize(a)

	// Run application on given PORT.
	// Default PORT equals to 3000.
	log.Fatal(a.Listen(config.PORT))

	// Close database connection.
	defer database.Disconnect()
}
