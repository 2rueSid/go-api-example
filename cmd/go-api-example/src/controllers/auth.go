// Controllers contained defined controllers.
package controllers

import (
	"errors"
	"time"

	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/config"
	tokenModel "github.com/2rueSid/go-api/cmd/go-api-example/src/model/token"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/model/user"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/types"
	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

// SignUp responsible for creating user.
func SignUp(ctx *fiber.Ctx) error {
	b := new(types.UserInput)
	c := make(chan *types.UserOutput)

	if err := ctx.BodyParser(b); err != nil {
		return fiber.ErrBadRequest
	}

	go user.Create(b, c)

	r := <-c

	err, s, u := r.Err, r.ErrStatus, r.User

	if err != nil {
		return fiber.NewError(s, err.Error())
	}

	authorized, err := generateAuthToken(u)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(authorized)
}

// SignIn responsible for sign in user.
func SignIn(ctx *fiber.Ctx) error {
	b := new(types.UserInput)
	c := make(chan *types.UserOutput)

	if err := ctx.BodyParser(b); err != nil {
		return fiber.ErrBadRequest
	}

	go user.SignIn(b, c)

	r := <-c

	err, s, u := r.Err, r.ErrStatus, r.User

	if err != nil {
		return fiber.NewError(s, err.Error())
	}

	authorized, err := generateAuthToken(u)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(authorized)
}

// Generates and save JWT token to the database.
// Also, return generated token, which is send to the frontend.
func generateAuthToken(u *db.UserModel) (*types.AuthorizedUser, error) {
	c := make(chan *types.TokenOutput)

	e := time.Now().Add(time.Hour * 72).Unix()

	claims := &types.JwtUserClaims{
		CurrentUser: types.CurrentUser{Name: u.Username, Email: u.Email, Id: u.ID},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: e,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := t.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return nil, errors.New(utils.StatusMessage(500))
	}

	go tokenModel.Create(
		&types.Token{UserId: u.ID, Token: s, Expiration: e},
		c,
	)

	if r := <-c; r.Err != nil {
		return nil, errors.New(utils.StatusMessage(500))
	}

	return &types.AuthorizedUser{Token: s}, nil
}
