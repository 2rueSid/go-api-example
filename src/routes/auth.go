// package in which defined routes
package routes

import (
	"github.com/2rueSid/go-api-example/src/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Function to define routes which applies to authorization functionality
func AuthRoutes(app *fiber.App) {
	route := app.Group("/auth", logger.New())

	route.Post("/signup", controllers.SignUp)
	route.Post("/signin", controllers.SignIn)
}
