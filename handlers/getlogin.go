package handlers

import (
	"app/views"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	return Render(c, views.Login())
}
