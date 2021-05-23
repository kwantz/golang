package main

import (
	"github.com/kwantz/golang/internal/app/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	handler.Routes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
