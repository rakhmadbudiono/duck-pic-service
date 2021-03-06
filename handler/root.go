package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rakhmadbudiono/duck-pic-service/config"
)

var cfg *config.Config

func init() {
	cfg = config.New()
}

// RegisterEndpoints register endpoint and its handler
func RegisterEndpoints(e *echo.Echo) {
	e.GET("/", Hello)
	e.GET("/image/static", FirstImage)
	e.GET("/image/random", RandomImage)
	e.GET("/image/:id", GetImage)
}
