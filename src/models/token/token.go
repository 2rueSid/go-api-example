package token

import (
	"github.com/2rueSid/go-api-example/src/config/database"
	"github.com/2rueSid/go-api-example/src/types"
)

var (
	client = database.Connect()
)

func Create(user *types.UserInput, chanel chan<- *types.UserOutput) {

}
