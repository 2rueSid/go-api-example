// middleware contains defined middlewares.
package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// CurrentUser used in protected routes.
// If there is a JWT token and it's valid,
// add CurrentUser struct instance into the application context.
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
