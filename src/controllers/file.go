// Package in which controllers are initialized
package controllers

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	"github.com/2rueSid/go-api-example/prisma/db"
	fileModel "github.com/2rueSid/go-api-example/src/models/file"
	"github.com/2rueSid/go-api-example/src/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

const (
	// Key that equals to key in multipart form data request field
	UPLOAD_FILES_KEY = "files"
)

// Function that save uploaded locally
// And return slice of saved files
func UploadFiles(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()

	userId := int(ctx.Locals("CurrentUser").(float64))

	fileChanel := make(chan *types.FileOutput)

	if err != nil {
		return fiber.NewError(400, utils.StatusMessage(400))
	}

	files := form.File[UPLOAD_FILES_KEY]
	var uploadedFiles []db.FileModel

	for _, file := range files {
		name, extension := generateName(file.Filename), getExtension(file.Filename)

		fileInput := &types.FileInput{
			Size:         int(file.Size),
			Name:         name,
			Extension:    extension,
			Originalname: file.Filename,
			UserId:       userId,
		}

		// Saved file into db
		go fileModel.Create(fileInput, fileChanel)

		savedFile := <-fileChanel

		if savedFile.Err != nil {
			return fiber.NewError(savedFile.Status, savedFile.Err.Error())
		}

		uploadedFiles = append(uploadedFiles, *savedFile.File)

		ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", name+extension))
	}

	return ctx.JSON(uploadedFiles)
}

func generateName(originalname string) string {
	var name bytes.Buffer

	name.WriteString(originalname)
	name.WriteString("_")
	name.WriteString(time.Now().String())

	return name.String()
}

func getExtension(filename string) string {
	extension := filepath.Ext(filename)

	return extension
}
