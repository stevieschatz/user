package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", CreateUser)

	// e.GET("/users/:id", GetUser)
	// e.PUT("/users/:id", UpdateUser)
	// e.DELETE("/users/:id", DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
