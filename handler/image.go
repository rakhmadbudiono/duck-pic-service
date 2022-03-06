package handler

import (
	"fmt"
	"net/http"

	"github.com/rakhmadbudiono/duck-pic-service/package/do"

	"github.com/labstack/echo/v4"
)

// FirstImage handle request for endpoint "/first"
func FirstImage(c echo.Context) error {
	return c.File("static/image.jpg")
}

// GetImage handler request for endpoint "/image/:id"
func GetImage(c echo.Context) error {
	id := c.Param("id")

	path := fmt.Sprintf("%s/%s.jpg", cfg.Duck.SpaceFolder, id)
	bytes, mimeType := do.GetObject(path)

	return c.Blob(http.StatusOK, mimeType, bytes)
}
