// package in which defined routes
package routes

import (
	"github.com/2rueSid/go-api-example/src/controllers"
	"github.com/2rueSid/go-api-example/src/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Protected route
func ProtectedRoutes(app fiber.Router) {
	route := app.Group("/protected", logger.New())

	route.Get("/", middleware.Protected(), controllers.ProtectedRoute)
}
