// Package in which controllers are initialized
package controllers

import (
	"github.com/2rueSid/go-api-example/src/models"
	"github.com/2rueSid/go-api-example/src/types"

	"github.com/gofiber/fiber/v2"
)

// Controller that responsible for creating user
func SignUp(ctx *fiber.Ctx) error {
	body := new(types.CreateUser)

	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := models.CreateUser(body)

	if err != nil {
		return fiber.ErrConflict
	}

	return ctx.JSON(user)
}

// Controller that responsible for sign in user
func SignIn(ctx *fiber.Ctx) error {

	return nil
}
