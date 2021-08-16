// Middleware contains defined middlewares.
package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// CurrentUser used in protected routes.
// If there is a JWT token and it's valid,
// add CurrentUser struct instance into the application context.
func CurrentUser(c *fiber.Ctx) error {
	t := c.Locals("user")

	u := t.(*jwt.Token).Claims.(jwt.MapClaims)

	if err := u.Valid(); err != nil {
		return errors.New("not valid")
	}

	id := u["id"]

	if id == nil {
		return errors.New("not valid")
	}

	c.Locals("CurrentUser", id)

	return c.Next()
}
