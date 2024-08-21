package handlers

import (
	"app/db"
	"app/models"
	"app/views"
	"database/sql"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c echo.Context) error {
	return Render(c, views.Login())
}

func PostLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := authenticateUser(email, password)
	if user == nil || err != nil {
		log.Error().Err(err).Msg("Authentication failed. Maybe a user is nil after authentication")
		return c.Redirect(http.StatusSeeOther, "/login?error=invalid_credentials")
	}

	sess, err := session.Get("session", c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
	}
	sess.Values["user_id"] = user.ID
	sess.Values["username"] = user.Username
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/account")
}

func authenticateUser(email, password string) (*models.User, error) {
	var user models.User
	stmt := `SELECT id, username, email, hashed_pwd FROM users WHERE email = $1`
	err := db.Db.QueryRow(stmt, email).Scan(&user.ID, &user.Username, &user.Email, &user.Hashed_pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msg("no user found")
			return nil, err
		}
		log.Error().Err(err).Msg("Database query failed")
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hashed_pwd), []byte(password))
	if err != nil {
		log.Error().Msg("Invalid password")
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return &user, nil
}
