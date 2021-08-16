// package to implement types
package types

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	"github.com/golang-jwt/jwt"
)

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

// Output that has user instance and generated JWT token,
// that passed to the frontend to implement JWT session
type AuthorizedUser struct {
	Token string `json:"token"`
}

// Current user payload that used within the application context
type CurrentUser struct {
	Name  string `json:"username"`
	Email string `json:"email"`
	Id    int    `json:"id"`
}

// JWT claims type
type JwtUserClaims struct {
	CurrentUser
	jwt.StandardClaims
}
