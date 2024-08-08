package main

import (
	"app/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", handlers.HelloHandler)
	e.GET("login", handlers.LoginHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
