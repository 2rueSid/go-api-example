package types

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
)

// Required token data, that need to create token and save it into the userTokens table
type Token struct {
	Expiration int64  `json:"expiration"`
	Token      string `json:"token"`
	UserId     int    `json:"user_id"`
}

// Create token function output
type TokenOutput struct {
	Token     *db.UserTokensModel
	Err       error
	ErrStatus int
}
