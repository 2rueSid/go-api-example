// types contains implemented types.
package types

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
)

// Token contains a required token data,
// that need to create token and save it into the user_tokens table.
type Token struct {
	Expiration int64  `json:"expiration"`
	Token      string `json:"token"`
	UserId     int    `json:"user_id"`
}

// TokenOutput describes create token function output.
type TokenOutput struct {
	Token     *db.UserTokensModel
	Err       error
	ErrStatus int
}
