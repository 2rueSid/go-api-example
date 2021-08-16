// routes contains defined routes.
package routes

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/src/controllers"
	"github.com/gofiber/fiber/v2"
)

// AuthRoutes is used to define routes,
// which are applies to the authorization functionality.
func AuthRoutes(app *fiber.App) {
	route := app.Group("/auth")

	route.Post("/signup", controllers.SignUp)
	route.Post("/signin", controllers.SignIn)
}
