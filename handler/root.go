package handler

import "github.com/labstack/echo/v4"

// RegisterEndpoints register endpoint and its handler
func RegisterEndpoints(e *echo.Echo) {
	e.GET("/", Hello)
	e.GET("/first", FirstImage)
}
