package repository

import (
	"github.com/deadlorpa/auth-app/interfaces"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization interfaces.AuthorizationRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
