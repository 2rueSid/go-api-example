// Package main implements web server.
// Here is the server starts
package main

import (
	"context"

	"github.com/2rueSid/go-api-example/src/types"

	"github.com/2rueSid/go-api-example/prisma/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/users", CreateUser)

	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := Users()
		if err != nil {
			c.SendStatus(500)
		}
		return c.JSON(users)
	})

	app.Listen(":3000")
}

// Get users by condition
// Where deleted_at field is nil
func Users() ([]db.UserModel, error) {
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	users, err := client.User.FindMany(db.User.DeletedAt.IsNull()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

/*
	Create new user.

	@body
	username - string
	email - string
    password - string
*/
func CreateUser(c *fiber.Ctx) error {

	body := new(types.CreateUser)

	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	createdUser, err := client.User.CreateOne(
		db.User.Email.Set(body.Email),
		db.User.Username.Set(body.Username),
		db.User.Password.Set(body.Password)).Exec(ctx)

	if err != nil {
		return fiber.ErrConflict
	}

	return c.JSON(createdUser)
}
