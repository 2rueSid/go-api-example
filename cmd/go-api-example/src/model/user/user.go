// User is used to define methods that are applied to user table.
package user

import (
	"errors"

	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/config/database"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/types"
	"github.com/gofiber/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	// Default cost, that used to hash password within the bcrypt library.
	DEFAULT_COST = 10
)

var (
	// Get database connection instance.
	client = database.Connect()
)

// Create creates user, and save it to the users table.
func Create(u *types.UserInput, c chan<- *types.UserOutput) {
	p, err := hashPassword(u.Password)

	if err != nil {
		c <- output(nil, 500, errors.New(utils.StatusMessage(500)))
	}

	created, err := client.User.CreateOne(
		db.User.Email.Set(u.Email),
		db.User.Username.Set(u.Username),
		db.User.Password.Set(p)).Exec(database.Context)

	if err != nil {
		c <- output(nil, 403, errors.New(utils.StatusMessage(403)))
	}

	c <- output(created, 0, nil)
}

// SignIn accepts user data, and return user if it's exists or return an error.
func SignIn(d *types.UserInput, c chan<- *types.UserOutput) {
	u, err := client.User.FindUnique(
		db.User.EmailUsername(db.User.Email.Equals(d.Email),
			db.User.Username.Equals(d.Username)),
	).Exec(database.Context)

	if err != nil {
		c <- output(nil, 401, errors.New(utils.StatusMessage(404)))
	}

	if err := comparePasswords(u.Password, d.Password); err != nil {
		c <- output(nil, 401, errors.New(utils.StatusMessage(401)))
	}

	c <- output(u, 0, nil)
}

// Output creates default output object.
func output(u *db.UserModel, s int, err error) *types.UserOutput {
	return &types.UserOutput{User: u, ErrStatus: s, Err: err}
}

// Hash given password, and return a hash.
func hashPassword(p string) (password string, err error) {
	v, err := bcrypt.GenerateFromPassword([]byte(p), DEFAULT_COST)

	if err != nil {
		return "", errors.New("err")
	}

	return string(v), nil
}

// Compare passwords, and return error if they are not the same.
func comparePasswords(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("err")
	}

	return nil
}
