// Routes contains defined routes.
package routes

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/src/controllers"
	"github.com/gofiber/fiber/v2"
)

// AuthRoutes is used to define routes,
// which are applies to the authorization functionality.
func AuthRoutes(a *fiber.App) {
	r := a.Group("/auth")

	r.Post("/signup", controllers.SignUp)
	r.Post("/signin", controllers.SignIn)
}
