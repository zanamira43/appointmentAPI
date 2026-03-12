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
func ValidateSignatureImage(file *dto.SignatureFileImage) error {

	if file.SignatureFileUrl == "" {
		return errors.New("signature image url is required")
	}

	return nil
}

// image uploader function to local storage
func UploadSignatureFileImage(c echo.Context) error {
	// coming images from frontend
	file, err := c.FormFile("signature")
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
	err = os.MkdirAll("./uploads/signature/", os.ModePerm)
	if err != nil {
		return err
	}

	// create file name with path as a destination
	dst, err := os.Create(fmt.Sprintf("./uploads/signature/%s", file.Filename))
	if err != nil {
		return err
	}

	// postpone closing file
	defer dst.Close()

	// copy the soruce file content  into destination file as content
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	// Determine scheme (http/https)
	scheme := c.Scheme() // uses Echo's detection; works with TLS or forwarded headers if middleware is set
	if scheme == "" {
		// Fallbacks if needed
		req := c.Request()
		if xfProto := req.Header.Get("X-Forwarded-Proto"); xfProto != "" {
			scheme = xfProto
		} else if req.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}

	host := c.Request().Host
	signatureImageURL := fmt.Sprintf("%s://%s/api/signature/image/%s", scheme, host, file.Filename)

	return c.JSON(http.StatusOK, echo.Map{
		"signature_file_url": signatureImageURL,
	})
}

// this is a function for deleting images in the storage dir
func DeleteSignatureImageFromStorage(url string) error {
	files := strings.SplitAfter(url, "/api/signature/image/")
	filename := files[1]

	if _, err := os.Stat("./uploads/signature/" + filename); err == nil {
		e := os.Remove("./uploads/signature/" + filename)
		if e != nil {
			log.Fatal(e)
		}
	}

	return nil
}
