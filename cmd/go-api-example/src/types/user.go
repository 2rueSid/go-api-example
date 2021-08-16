// types contains implemented types.
package types

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	"github.com/golang-jwt/jwt"
)

// UserInput describes which data
// should pass while creating a new user.
type UserInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserOutput describes the CreateUser channel type.
type UserOutput struct {
	User      *db.UserModel
	Err       error
	ErrStatus int
}

// AuthorizedUser describes an output of the controllers
// that applies to the user' session, and contains a JWT token, which
// will be used in the frontend to implement the JWT session.
type AuthorizedUser struct {
	Token string `json:"token"`
}

// CurrentUser describes the current user
// which we could get from the JWT payload.
type CurrentUser struct {
	Name  string `json:"username"`
	Email string `json:"email"`
	Id    int    `json:"id"`
}

// JwtUserClaims describes JWT claims.
// Those claims are used as the JWT token body.
type JwtUserClaims struct {
	CurrentUser
	jwt.StandardClaims
}
