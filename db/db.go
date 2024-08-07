package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var db *sql.DB

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	getConnection()
}

func getConnection() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	connection := os.Getenv("POSTGRES_CONNECTION")
	db, err = sql.Open("postgres", connection)
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
			password TEXT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		)`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create users table")
	}
	log.Info().Msg("Users table created")

}
