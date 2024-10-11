package helpers

import (
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func ImageUploader(c *gin.Context) (string, error) {
	file, err := c.FormFile("image")

	if err != nil {
		return "", errors.New("Failed to upload image111")
	}

	err = os.Mkdir("assets/images", 0775)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	filePath := "assets/images/" + file.Filename

	err = c.SaveUploadedFile(file, filePath)

	if err != nil {
		return "", errors.New("Failed to upload image")
	}

	return filePath, nil
}
