package handlers

import (
	"app/db"
	"app/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c echo.Context) error {
	return Render(c, views.Register())
}

// Can add SQLBuilder(squirrel) or GORM instead.
func RegisterUser(username, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return err
	}

	stmt := `
		INSERT INTO users (username, email, hashed_pwd)
		VALUES ($1, $2, $3)`
	db.Db.Exec(stmt, username, email, string(hashedPassword))
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
