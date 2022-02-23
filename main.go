package main

import (
	"fmt"

	"github.com/rakhmadbudiono/duck-pic-service/config"
	"github.com/rakhmadbudiono/duck-pic-service/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.New()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Hello)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	e.Logger.Fatal(e.Start(addr))
}
