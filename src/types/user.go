// package to implement types
package types

import "github.com/2rueSid/go-api-example/prisma/db"

// DTO that describes which data
// should pass while createing new user
type UserInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Struct which use to describe CreateUser chanel type
type UserOutput struct {
	User      *db.UserModel
	Err       error
	ErrStatus int
}

type AuthorizedUser struct {
	User  *db.UserModel `json:"user"`
	Token string        `json:"token"`
}
