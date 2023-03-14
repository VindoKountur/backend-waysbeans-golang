package middleware

import (
	dto "backendwaysbeans/dto/result"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("photo")
		method := c.Request().Method

		if err != nil {
			if method == "PATCH" && err.Error() == "http: no such file" {
				c.Set("dataFile", "")
				return next(c)
			}
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		defer tempFile.Close()

		if _, err = io.Copy(tempFile, src); err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		data := tempFile.Name()

		c.Set("dataFile", data)
		return next(c)
	}
}
