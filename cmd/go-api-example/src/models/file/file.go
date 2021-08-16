// file is used to define methods that are applied to file table.
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
func Create(file *types.FileInput, chanel chan<- *types.FileOutput) {
	download := "/uploads/" + file.Name
	createdFile, err := client.File.CreateOne(
		db.File.Name.Set(file.Name),
		db.File.Originalname.Set(file.Originalname),
		db.File.Size.Set(file.Size),
		db.File.Extension.Set(file.Extension),
		db.File.User.Link(db.User.ID.Equals(file.UserId)),
		db.File.Download.Set(download),
	).Exec(database.Context)

	if err != nil {
		result := &types.FileOutput{
			createdFile,
			types.ErrorOutput{Err: errors.New("file not created"), Status: 500},
		}

		chanel <- result
		return
	}

	chanel <- &types.FileOutput{createdFile, types.ErrorOutput{Err: nil, Status: 0}}
}
