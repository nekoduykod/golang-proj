package models

type User struct {
	ID         int
	Username   string
	Email      string
	Hashed_pwd string
}
