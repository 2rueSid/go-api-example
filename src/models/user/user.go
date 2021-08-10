// Package in which database queries are initialized
package user

import (
	"errors"

	"github.com/2rueSid/go-api-example/prisma/db"
	"github.com/2rueSid/go-api-example/src/config/database"
	"github.com/2rueSid/go-api-example/src/types"
	"golang.org/x/crypto/bcrypt"
)

var (
	DEFAULT_COST = 10
)

// Function that creates user, and save it to the users table
func Create(user *types.CreateUser) (*db.UserModel, error) {
	client := database.Connect()

	password, err := hashPassword(user.Password)

	if err != nil {
		return nil, errors.New("error while trying to hash password")
	}

	createdUser, err := client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Username.Set(user.Username),
		db.User.Password.Set(password)).Exec(database.Context)

	if err != nil {
		return nil, errors.New("user not created")
	}

	return createdUser, nil
}

// Hash given password, and return a hash
func hashPassword(password string) (string, error) {
	value, err := bcrypt.GenerateFromPassword([]byte(password), DEFAULT_COST)

	if err != nil {
		return "", nil
	}

	return string(value), nil
}

// func comparePasswords(hashed string, password string) error {
// 	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
// 		return errors.New("passwords not same")
// 	}

// 	return nil
// }
