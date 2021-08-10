// Package in which database queries are initialized
package models

import (
	"errors"

	"github.com/2rueSid/go-api-example/prisma/db"
	"github.com/2rueSid/go-api-example/src/config"
	"github.com/2rueSid/go-api-example/src/types"
)

// Function that creates user, and save it to the users table
func CreateUser(user *types.CreateUser) (*db.UserModel, error) {
	client := config.Connect()

	createdUser, err := client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Username.Set(user.Username),
		db.User.Password.Set(user.Password)).Exec(config.Context)

	if err != nil {
		return nil, errors.New("user not created")
	}

	config.Disconnect()

	return createdUser, nil
}
