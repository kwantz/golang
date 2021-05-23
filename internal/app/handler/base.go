package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
