// Controllers contained defined controllers.
package controllers

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
	fileModel "github.com/2rueSid/go-api/cmd/go-api-example/src/model/file"
	"github.com/2rueSid/go-api/cmd/go-api-example/src/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

const (
	// Key that equals to key in multipart form data request field.
	UPLOAD_FILES_KEY = "files"
)

// UploadFiles save uploaded locally,
// and return slice of saved files.
func UploadFiles(c *fiber.Ctx) error {
	f, err := c.MultipartForm()
	userID := int(c.Locals("CurrentUser").(float64))

	fm := &fileModel.File{}

	fileChanel := make(chan *types.FileOutput)

	if err != nil {
		return fiber.NewError(400, utils.StatusMessage(400))
	}

	files := f.File[UPLOAD_FILES_KEY]
	var uploaded []db.FileModel

	for _, file := range files {
		n, e := generateName(file.Filename), getExtension(file.Filename)

		fileInput := &types.FileInput{
			Size:         int(file.Size),
			Name:         n,
			Extension:    e,
			Originalname: file.Filename,
			UserId:       userID,
		}

		// Saved file into db.
		go fm.Create(fileInput, fileChanel)

		saved := <-fileChanel

		if saved.Err != nil {
			return fiber.NewError(saved.Status, saved.Err.Error())
		}

		uploaded = append(uploaded, *saved.File)

		err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", n+e))

		if err != nil {
			return fiber.NewError(saved.Status, saved.Err.Error())
		}
	}

	return c.JSON(uploaded)
}

// generateName generate name of given file, and return generated name.
func generateName(n string) string {
	var name bytes.Buffer

	name.WriteString(n)
	name.WriteString("_")
	name.WriteString(time.Now().String())

	return name.String()
}

// getExtension is used to get extension of given file.
func getExtension(n string) string {
	e := filepath.Ext(n)

	return e
}
