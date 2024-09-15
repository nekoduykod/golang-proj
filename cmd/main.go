package main

import (
	"app/internal/handlers"
	"app/pkg/repository"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var store *sessions.CookieStore

func main() {
	e := echo.New()

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	})

	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Failed to load env file")
	}

	database := repository.GetConnection()
	userRepo := repository.NewUserRepository(database)

	h := handlers.NewHandlers(userRepo)

	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal().Msg("SESSION SECRET env is not set")
	}
	store = sessions.NewCookieStore([]byte(sessionSecret))
	e.Use(session.Middleware(store))

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/register")
	})

	e.GET("register", h.RegisterHandler)
	e.GET("login", h.LoginHandler)
	e.GET("account", h.AccountHandler)

	e.POST("register", h.PostRegister)
	e.POST("login", h.PostLogin)
	e.POST("logout", h.LogoutHandler)

	// http://localhost:8080/
	e.Logger.Fatal(e.Start(":8080"))
}
