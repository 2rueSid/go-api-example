// Package in which controllers are initialized
package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func ProtectedRoute(ctx *fiber.Ctx) error {

	return ctx.JSON("id: 22")
}
