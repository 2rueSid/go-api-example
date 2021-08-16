// File is used to define methods that are applied to file table.
package file

import (
	"errors"

	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/config/database"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/types"
)

var (
	// Get database connection instance.
	client = database.Client
)

// Add file to the database.
func Create(f *types.FileInput, c chan<- *types.FileOutput) {
	d := "/uploads/" + f.Name
	created, err := client.File.CreateOne(
		db.File.Name.Set(f.Name),
		db.File.Originalname.Set(f.Originalname),
		db.File.Size.Set(f.Size),
		db.File.Extension.Set(f.Extension),
		db.File.User.Link(db.User.ID.Equals(f.UserId)),
		db.File.Download.Set(d),
	).Exec(database.Context)

	if err != nil {
		c <- &types.FileOutput{
			File:      created,
			ErrOutput: types.ErrOutput{Err: errors.New("file not created"), Status: 500},
		}

		return
	}

	c <- &types.FileOutput{File: created, ErrOutput: types.ErrOutput{Err: nil, Status: 0}}
}
