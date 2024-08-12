package handlers

import (
	"app/db"
	"app/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func RegisterHandler(c echo.Context) error {
	return Render(c, views.Register())
}

// Can add SQLBuilder(squirrel) or GORM instead.
func RegisterUser(username, email, password string) error {
	stmt := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)`
	db.Db.Exec(stmt, username, email, password)
	return nil
}

func PostRegister(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := RegisterUser(username, email, password)
	if err != nil {
		log.Error().Err(err).Msg("Failed to register new user")
		return c.Redirect(http.StatusSeeOther, "/register")
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
