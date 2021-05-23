package usecase

import (
	"strconv"

	"github.com/kwantz/golang/internal/entity"
	"github.com/labstack/echo/v4"
)

var (
	users = map[int]*entity.User{}
	seq   = 1
)

func CreateUser(c echo.Context) (*entity.User, error) {
	u := &entity.User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return nil, err
	}
	users[u.ID] = u
	seq++

	return u, nil
}

func GetUser(c echo.Context) (*entity.User, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	return users[id], nil
}

func UpdateUser(c echo.Context) (*entity.User, error) {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		return nil, err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return users[id], nil
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return nil
}

func GetAllUsers(c echo.Context) (map[int]*entity.User, error) {
	return users, nil
}
