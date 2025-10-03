package helpers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/zanamira43/appointment-api/dto"
)

// validate image url
func ValidateImage(img *dto.Image) error {

	if img.PatientImageUrl == "" {
		return errors.New("image url is required")
	}

	return nil
}

// image uploader function to local storage
func UploadImage(c echo.Context) error {
	// coming images from frontend
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// open image file
	src, err := file.Open()
	if err != nil {
		return err
	}
	// postpone closing file for the end of process
	defer src.Close()

	// create uploads directory if not exists
	// if the directory exists already it does nothing
	err = os.MkdirAll("./uploads/", os.ModePerm)
	if err != nil {
		return err
	}

	// create file name with path as a destination
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
	if err != nil {
		return err
	}

	// postpone closing file
	defer dst.Close()

	// copy the soruce file content  into destination file as content
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"patient_image_url": c.Request().Host + "/api/image/" + file.Filename,
	})
}

// this is a function for deleting images in the storage dir
func DeleteImageFromStorage(url string) error {
	files := strings.SplitAfter(url, "/api/image/")
	filename := files[1]

	if _, err := os.Stat("./uploads/" + filename); err == nil {
		e := os.Remove("./uploads/" + filename)
		if e != nil {
			log.Fatal(e)
		}
	}

	return nil
}
