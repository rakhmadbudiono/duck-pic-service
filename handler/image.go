package handler

import (
	"github.com/labstack/echo/v4"
)

// FirstImage handle request for endpoint "/first"
func FirstImage(c echo.Context) error {
	return c.File("static/image.jpg")
}
