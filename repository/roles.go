package repository

import (
	"fmt"

	"github.com/deadlorpa/auth-app/model"
	"github.com/jmoiron/sqlx"
)

type RolesRepository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) *RolesRepository {
	return &RolesRepository{db: db}
}

func (r *RolesRepository) GetById(id string) (role model.Role, err error) {
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id=$1", roleTable)
	err = r.db.Get(&role, query, id)

	return role, err
}
