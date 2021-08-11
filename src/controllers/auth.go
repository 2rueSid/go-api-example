// Package in which controllers are initialized
package controllers

import (
	"time"

	"github.com/2rueSid/go-api-example/src/config"
	"github.com/2rueSid/go-api-example/src/models/user"
	"github.com/2rueSid/go-api-example/src/types"
	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
)

// Controller that responsible for creating user
func SignUp(ctx *fiber.Ctx) error {
	body := new(types.UserInput)
	c := make(chan *types.UserOutput)

	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	go user.Create(body, c)

	result := <-c

	err, status, user := result.Err, result.ErrStatus, result.User

	if err != nil {
		return fiber.NewError(status, err.Error())
	}
	return ctx.JSON(user)
}

// Controller that responsible for sign in user
func SignIn(ctx *fiber.Ctx) error {
	body := new(types.UserInput)
	c := make(chan *types.UserOutput)

	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	go user.SignIn(body, c)

	result := <-c

	err, status, user := result.Err, result.ErrStatus, result.User

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expiration := time.Now().Add(time.Hour * 72).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = expiration

	signed, err := token.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return fiber.NewError(500, "server error")
	}

	authorizedUser := &types.AuthorizedUser{User: user, Token: signed}

	return ctx.JSON(authorizedUser)
}
