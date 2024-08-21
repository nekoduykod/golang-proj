package main

import (
	"app/db"
	"app/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

var store *sessions.CookieStore

func main() {
	e := echo.New()
	db.DB_init()

	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Failed to load env file")
	}

	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal().Msg("SESSION SECRET env is not set")
	}
	store = sessions.NewCookieStore([]byte(sessionSecret))
	e.Use(session.Middleware(store))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/register")
	})

	e.GET("register", handlers.RegisterHandler)
	e.GET("login", handlers.LoginHandler)
	e.GET("account", handlers.AccountHandler)

	e.POST("register", handlers.PostRegister)
	e.POST("login", handlers.PostLogin)
	e.POST("logout", handlers.LogoutHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
