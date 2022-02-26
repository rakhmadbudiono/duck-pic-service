package handler_test

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func TestMain(m *testing.M) {
	e = echo.New()

	os.Exit(m.Run())
}
