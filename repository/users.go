package repository

import (
	"fmt"

	"github.com/deadlorpa/auth-app/model"
	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) Create(user model.User) (id string, err error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *UsersRepository) GetBySignIn(userSignIn model.UserSignInRequest) (user model.User, err error) {
	query := fmt.Sprintf("SELECT id, username, id_role AS IdRole FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err = r.db.Get(&user, query, userSignIn.Username, userSignIn.Password)

	return user, err
}
