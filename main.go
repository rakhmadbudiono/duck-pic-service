package main

import (
	"fmt"
	"log"

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

	handler.RegisterEndpoints(e)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := e.Start(addr); err != nil {
		log.Fatalf("Couldn't start REST server: %s", err)
	}
}
