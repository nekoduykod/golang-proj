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
)

// func SecureRandStringBytes(n int) (string, error) {
// 	b := make([]byte, n)
// 	if _, err := rand.Read(b); err != nil {
// 		return "", err
// 	}
// 	return base64.URLEncoding.EncodeToString(b)[:n], nil
// }
// secret := []byte(SecureRandStringBytes(20))

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
	stmt := `SELECT id, username, email, password FROM users WHERE email = $1`
	err := db.Db.QueryRow(stmt, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msg("no user found")
			return nil, err
		}
		log.Error().Err(err).Msg("Database query failed")
		return nil, err
	}

	if password != user.Password {
		log.Error().Msg("Invalid password") // eeeww what`s that
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return &user, nil
}
