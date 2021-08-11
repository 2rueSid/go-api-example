// Package in which database queries are initialized
package user

import (
	"errors"

	"github.com/2rueSid/go-api-example/prisma/db"
	"github.com/2rueSid/go-api-example/src/config/database"
	"github.com/2rueSid/go-api-example/src/types"
	"golang.org/x/crypto/bcrypt"
)

const (
	DEFAULT_COST = 10
)

var (
	client = database.Connect()
)

// Function that creates user, and save it to the users table
func Create(user *types.UserInput, chanel chan<- *types.UserOutput) {
	password, err := hashPassword(user.Password)

	if err != nil {
		result := &types.UserOutput{
			User:      nil,
			ErrStatus: 500,
			Err:       errors.New("error while trying to hash password"),
		}

		chanel <- result
		return
	}

	createdUser, err := client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Username.Set(user.Username),
		db.User.Password.Set(password)).Exec(database.Context)

	if err != nil {
		result := &types.UserOutput{
			User:      nil,
			ErrStatus: 500,
			Err:       errors.New("user not created"),
		}

		chanel <- result
		return
	}

	result := &types.UserOutput{User: createdUser, Err: nil, ErrStatus: 0}

	chanel <- result
}

// Function that accepts user data, and return user if it's exists or return an error
func SignIn(data *types.UserInput, chanel chan<- *types.UserOutput) {
	user, err := client.User.FindUnique(db.User.EmailUsername(db.User.Email.Equals(data.Email), db.User.Username.Equals(data.Username))).Exec(database.Context)

	if err != nil {
		result := &types.UserOutput{
			User:      nil,
			ErrStatus: 404,
			Err:       errors.New("user not exists"),
		}

		chanel <- result
	}

	if err := comparePasswords(user.Password, data.Password); err != nil {
		result := &types.UserOutput{
			User:      nil,
			ErrStatus: 403,
			Err:       errors.New("passwords are not the same"),
		}

		chanel <- result
	}

	chanel <- &types.UserOutput{User: user, Err: nil}
}

// Hash given password, and return a hash
func hashPassword(password string) (string, error) {
	value, err := bcrypt.GenerateFromPassword([]byte(password), DEFAULT_COST)

	if err != nil {
		return "", nil
	}

	return string(value), nil
}

// Compare passwords, and return error if they are not the same
func comparePasswords(hashed string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return errors.New("passwords not same")
	}

	return nil
}