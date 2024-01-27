package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/deadlorpa/auth-app/configs"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	roleTable  = "roles"
)

func NewPostgresDB(cfg configs.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(cfg configs.DBConfig) {
	if m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)); err != nil {
		log.Fatalf("!!! cannot migrate db: %s", err.Error())
	} else {
		if err = m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Printf("~~~ no migration changes")
			} else {
				log.Fatalf("!!! cannot migrate db: %s", err.Error())
			}
		}
	}
}
