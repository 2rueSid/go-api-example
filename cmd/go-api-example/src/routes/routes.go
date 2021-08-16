// Routes contains defined routes.
package routes

import "github.com/gofiber/fiber/v2"

// Initialize used to initialize all routes that were defined in the routes directory.
func Initialize(a *fiber.App) {
	AuthRoutes(a)
	UploadRoutes(a)
}
