package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakhmadbudiono/duck-pic-service/handler"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// expected result
	expectedStatusCode := http.StatusOK
	expectedResBody := "Hello world!"

	e.GET("/", handler.Hello)

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, expectedStatusCode, rec.Code, fmt.Sprintf("TestHello: status code should be %d, but got %d", expectedStatusCode, rec.Code))
	assert.Equal(t, expectedResBody, rec.Body.String(), fmt.Sprintf("TestHello: response body should be %s, but got %s", expectedResBody, rec.Body.String()))
}
