package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Db *sql.DB

func DB_init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	getConnection()
}

func getConnection() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load end")
	}

	connection := os.Getenv("POSTGRES_CONNECTION")
	Db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	log.Info().Msg("Connection initialized")

	createMigration()
}

func createMigration() {
	stmt := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			hashed_pwd TEXT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW()
		)`

	_, err := Db.Exec(stmt)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create users table")
	}
	log.Info().Msg("Users table created")

}
