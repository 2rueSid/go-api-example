// Package main implements web server.
// Here is the server starts
package main

import (
	"github.com/2rueSid/go-api-example/src/config"
	"github.com/2rueSid/go-api-example/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Initialize(*app)

	routes.Initialize(app)

	app.Listen(config.PORT)
}
