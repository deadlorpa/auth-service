package repository

import (
	"github.com/deadlorpa/auth-app/interfaces"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users interfaces.UsersRepository
	Roles interfaces.RolesRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUsersRepository(db),
		Roles: NewRolesRepository(db),
	}
}
