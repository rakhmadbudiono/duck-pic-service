package handler

import "github.com/labstack/echo/v4"

func RegisterEndpoints(e *echo.Echo) {
	e.GET("/", Hello)
}
