package repository

import (
	"app/internal/models"
	"database/sql"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	RegisterUser(username, email, password string) error
	GetUserByEmail(email string) (*models.User, error)
	VerifyPassword(hashedPassword, password string) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &repo{db: db}
}

type repo struct {
	db *sql.DB
}

// Can add SQLBuilder(squirrel) or GORM instead.
func (db *repo) RegisterUser(username, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Failed to hash password")
		return err
	}

	stmt := `
		INSERT INTO users (username, email, hashed_pwd)
		VALUES ($1, $2, $3)`
	_, err = db.db.Exec(stmt, username, email, string(hashedPassword))
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert new user")
		return err
	}
	return nil
}

func (db *repo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	stmt := `SELECT id, username, email, hashed_pwd FROM users WHERE email = $1`
	err := db.db.QueryRow(stmt, email).Scan(&user.ID, &user.Username, &user.Email, &user.Hashed_pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msg("No user found")
			return nil, err
		}
		log.Error().Err(err).Msg("Database query failed")
		return nil, err
	}
	return &user, nil
}

func (db *repo) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
