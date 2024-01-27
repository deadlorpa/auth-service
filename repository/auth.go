package repository

import (
	"fmt"

	"github.com/deadlorpa/auth-app/model"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user model.User) (id string, err error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(userSignIn model.UserSignInRequest) (user model.User, err error) {
	query := fmt.Sprintf("SELECT id, username, id_role AS IdRole FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err = r.db.Get(&user, query, userSignIn.Username, userSignIn.Password)

	return user, err
}

func (r *AuthRepository) GetRole(id string) (role model.Role, err error) {
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id=$1", roleTable)
	err = r.db.Get(&role, query, id)

	return role, err
}
