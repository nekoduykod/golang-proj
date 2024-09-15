package handlers

import (
	"app/internal/models"
	"app/views"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (h *Handlers) AccountHandler(c echo.Context) error {
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

	user := models.User{ID: userID, Username: username}
	return Render(c, views.Account(user))
}

func (h *Handlers) LogoutHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	log.Info().Msg("Logout")

	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/login")
}
