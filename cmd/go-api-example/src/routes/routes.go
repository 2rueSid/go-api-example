// routes contains defined routes.
package routes

import "github.com/gofiber/fiber/v2"

// Initialize used to initialize all routes that were defined in the routes directory.
func Initialize(app *fiber.App) {
	AuthRoutes(app)
	UploadRoutes(app)
}
