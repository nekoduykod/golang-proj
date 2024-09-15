package handlers

import (
	"app/pkg/repository"
)

type Handlers struct {
	repo repository.UserRepository
}

func NewHandlers(repo repository.UserRepository) *Handlers {
	return &Handlers{repo: repo}
}
