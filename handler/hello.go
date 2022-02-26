package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello handle request for root endpoint "/"
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}
