package handlers

import (
	"app/models"
	"app/views"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func AccountHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	userID, ok := sess.Values["user_id"].(int)

	if !ok || userID == 0 {
		log.Error().Msg("User ID not found in session or type assertion failed")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	username, ok := sess.Values["username"].(string)
	if !ok || username == "" {
		log.Error().Err(err).Msg("Username not found in session or type assertion failed")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	// var user models.User
	// err = db.Db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)
	// if err != nil {
	// 	log.Error().Err(err).Msg("failed to fetch user from database")
	// 	return c.Redirect(http.StatusSeeOther, "/login")
	// }

	user := models.User{ID: userID, Username: username} // Create user object from session values
	return Render(c, views.Account(user))
}

func LogoutHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/login")
}
