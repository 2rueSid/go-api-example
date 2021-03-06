// Middleware contains defined middlewares.
package middleware

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/src/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes.
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.JWT_SECRET),
		ErrorHandler: jwtError,
	})
}

// Error handler.
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "400", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "401", "message": "Invalid or expired JWT", "data": nil})
}
