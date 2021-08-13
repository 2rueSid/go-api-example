// package in which defined routes
package routes

import "github.com/gofiber/fiber/v2"

// Initialize all routes that were defined in the routes directory
func Initialize(app *fiber.App) {
	AuthRoutes(app)
	UploadRoutes(app)
}
