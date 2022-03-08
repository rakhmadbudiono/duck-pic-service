package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rakhmadbudiono/duck-pic-service/package/do"

	"github.com/labstack/echo/v4"
)

// FirstImage handle request for endpoint "/image/static"
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

// RandomImage handler request for endpoint "/image/random"
func RandomImage(c echo.Context) error {
	timestamp := time.Now()

	param := c.QueryParam("hour")
	hour, err := strconv.Atoi(param)
	if err != nil {
		log.Printf("Unable to convert hour query param to int %v", err)
	}

	h := time.Duration(hour) * time.Hour
	timestamp = timestamp.Truncate(h)

	seed := timestamp.Unix()
	id := seed%cfg.Duck.CountDuck + 1

	path := fmt.Sprintf("%s/%d.jpg", cfg.Duck.SpaceFolder, id)
	bytes, mimeType := do.GetObject(path)

	return c.Blob(http.StatusOK, mimeType, bytes)
}
