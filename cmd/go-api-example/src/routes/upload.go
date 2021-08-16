// package in which defined routes
package routes

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/src/controllers"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/middleware"
	"github.com/gofiber/fiber/v2"
)

// Function to define routes which applies to upload functionality
func UploadRoutes(app *fiber.App) {
	route := app.Group("/upload")

	route.Post("/files", middleware.Protected(), middleware.CurrentUser, controllers.UploadFiles)
}
