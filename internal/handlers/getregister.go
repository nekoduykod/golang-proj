package handlers

import (
	"app/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (h *Handlers) RegisterHandler(c echo.Context) error {
	return Render(c, views.Register())
}

func (h *Handlers) PostRegister(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := h.repo.RegisterUser(username, email, password)
	if err != nil {
		log.Error().Err(err).Msg("Failed to register new user")
		return c.Redirect(http.StatusSeeOther, "/register?error=registration_failed")
	}
	log.Info().Msg("Register success")

	return c.Redirect(http.StatusSeeOther, "/login")
}
