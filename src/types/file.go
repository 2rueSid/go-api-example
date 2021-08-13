package types

import (
	"github.com/2rueSid/go-api-example/prisma/db"
)

// Fields that need to save file into the db
type FileInput struct {
	Name         string `json:"name"`
	Originalname string `json:"originalname"`
	Size         int    `json:"size"`
	Extension    string `json:"extension"`
	UserId       int    `json:"user_id"`
}

// Default error output, used basically in models of application
type ErrorOutput struct {
	Err    error
	Status int
}

// File output, default return value form file model methods
type FileOutput struct {
	File *db.FileModel
	ErrorOutput
}
