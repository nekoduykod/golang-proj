package handlers

import (
	"app/views"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (h *Handlers) LoginHandler(c echo.Context) error {
	return Render(c, views.Login())
}

func (h *Handlers) PostLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := h.repo.GetUserByEmail(email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return c.Redirect(http.StatusSeeOther, "/login?error=invalid_credentials")
	}

	err = h.repo.VerifyPassword(user.Hashed_pwd, password)
	if err != nil {
		log.Error().Msg("Invalid password")
		return c.Redirect(http.StatusSeeOther, "/login?error=invalid_credentials")
	}

	sess, err := session.Get("session", c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
	}
	log.Info().Msg("Login success")

	sess.Values["user_id"] = user.ID
	sess.Values["username"] = user.Username
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "/account")
}
