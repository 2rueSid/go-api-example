package token

import (
	"errors"
	"time"

	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/config/database"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/types"
)

var (
	// Get database connection instance
	client = database.Connect()
)

// Save token into userTokens table
func Create(token *types.Token, chanel chan<- *types.TokenOutput) {
	createdToken, err := client.UserTokens.CreateOne(
		db.UserTokens.Lifetime.Set(time.Unix(token.Expiration, 0)),
		db.UserTokens.Token.Set(token.Token),
		db.UserTokens.User.Link(db.User.ID.Equals(token.UserId)),
	).Exec(database.Context)

	if err != nil {
		result := &types.TokenOutput{Err: errors.New("token not created"), ErrStatus: 500, Token: nil}

		chanel <- result
		return
	}

	chanel <- &types.TokenOutput{Err: nil, ErrStatus: 0, Token: createdToken}
}
