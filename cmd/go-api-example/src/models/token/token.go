// Token is used to define methods that are applied to user token.
package token

import (
	"errors"
	"time"

	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/config/database"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/types"
)

var (
	// Get database connection instance.
	client = database.Connect()
)

// Save token into userTokens table.
func Create(t *types.Token, c chan<- *types.TokenOutput) {
	created, err := client.UserTokens.CreateOne(
		db.UserTokens.Lifetime.Set(time.Unix(t.Expiration, 0)),
		db.UserTokens.Token.Set(t.Token),
		db.UserTokens.User.Link(db.User.ID.Equals(t.UserId)),
	).Exec(database.Context)

	if err != nil {
		r := &types.TokenOutput{Err: errors.New("token not created"), ErrStatus: 500, Token: nil}

		c <- r
		return
	}

	c <- &types.TokenOutput{Err: nil, ErrStatus: 0, Token: created}
}
