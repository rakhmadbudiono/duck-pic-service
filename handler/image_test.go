package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
)

func TestFirstImage(t *testing.T) {
	// expected result
	expectedStatusCode := http.StatusOK

	e.GET("/image/first", func(c echo.Context) error {
		return c.File("../static/image.jpg")
	})

	req, _ := http.NewRequest("GET", "/image/first", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, expectedStatusCode, rec.Code, fmt.Sprintf("TestFirstImage: status code should be %d, but got %d", expectedStatusCode, rec.Code))
}
