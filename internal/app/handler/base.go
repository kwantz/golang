package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kwantz/golang/internal/app/usecase"
)

func Routes(e *echo.Echo) {
	e.GET("/", hello)
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createUser(c echo.Context) error {
	u, err := usecase.CreateUser(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	u, err := usecase.GetUser(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func updateUser(c echo.Context) error {
	u, err := usecase.UpdateUser(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	err := usecase.DeleteUser(c)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	u, err := usecase.GetAllUsers(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
