package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	users = map[int]*User{}
	seq   = 1
)

// CreateUser by name
func CreateUser(c echo.Context) error {
	// u := &User{
	// 	ID: seq,
	// }
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }
	// users[u.ID] = u
	// seq++

	u := "teststevie"

	ConnectDataBase()

	return c.JSON(http.StatusCreated, u)
}

// // GetUser By ID
// func GetUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	return c.JSON(http.StatusOK, users[id])
// }

// // UpdateUser By ID
// func UpdateUser(c echo.Context) error {
// 	u := new(user)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	users[id].Name = u.Name
// 	return c.JSON(http.StatusOK, users[id])
// }

// // DeleteUser By ID
// func DeleteUser(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	delete(users, id)
// 	return c.NoContent(http.StatusNoContent)
// }
