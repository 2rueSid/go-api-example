// types contains implemented types.
package types

import (
	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
)

// FileInput has fields that need to save files into the DB.
type FileInput struct {
	Name         string `json:"name"`
	Originalname string `json:"originalname"`
	Size         int    `json:"size"`
	Extension    string `json:"extension"`
	UserId       int    `json:"user_id"`
}

// ErrorOutput default error output, used basically in models of application.
type ErrorOutput struct {
	Err    error
	Status int
}

// FileOutput, default return value form file model methods.
type FileOutput struct {
	File *db.FileModel
	ErrorOutput
}
