// Package in which controllers are initialized
package controllers

import (
	"errors"
	"time"

	"github.com/2rueSid/go-api-example/prisma/db"
	"github.com/2rueSid/go-api-example/src/config"
	tokenModel "github.com/2rueSid/go-api-example/src/models/token"
	"github.com/2rueSid/go-api-example/src/models/user"
	"github.com/2rueSid/go-api-example/src/types"
	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
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

	authorizedUser, err := generateSignInToken(user)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(authorizedUser)
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

	authorizedUser, err := generateSignInToken(user)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(authorizedUser)
}

// Generates and save JWT token to the database
// Also, return generated token, which is send to the frontend
func generateSignInToken(user *db.UserModel) (*types.AuthorizedUser, error) {
	tokenC := make(chan *types.TokenOutput)

	expiration := time.Now().Add(time.Hour * 72).Unix()

	claims := &types.JwtUserClaims{
		types.CurrentUser{Name: user.Username, Email: user.Email, Id: user.ID},
		jwt.StandardClaims{
			ExpiresAt: expiration,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return nil, errors.New(utils.StatusMessage(500))
	}

	go tokenModel.Create(
		&types.Token{UserId: user.ID, Token: signed, Expiration: expiration},
		tokenC,
	)

	if tokenResult := <-tokenC; tokenResult.Err != nil {
		return nil, errors.New(utils.StatusMessage(500))
	}

	return &types.AuthorizedUser{Token: signed}, nil
}
