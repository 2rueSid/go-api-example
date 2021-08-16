// Package in which middlewares are initialized
package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Middleware that used in protected routes
// If there is a JWT token and it's valid
// Add CurrentUser struct instance into the application context
func CurrentUser(ctx *fiber.Ctx) error {
	token := ctx.Locals("user")

	user := token.(*jwt.Token).Claims.(jwt.MapClaims)

	if err := user.Valid(); err != nil {
		return errors.New("not valid")
	}

	id := user["id"]

	if id == nil {
		return errors.New("not valid")
	}

	ctx.Locals("CurrentUser", id)

	return ctx.Next()
}
