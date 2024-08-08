package handlers

import (
	"app/views"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	return Render(c, views.Hello())
}
