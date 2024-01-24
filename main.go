package main

import (
	"errors"
	"log"

	"github.com/deadlorpa/auth-app/handler"
	"github.com/deadlorpa/auth-app/repository"
	"github.com/deadlorpa/auth-app/service"
	"github.com/deadlorpa/auth-app/structure"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("!!! cannot read configs: %s", err.Error())
	}
	MigrateDb()
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(structure.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("!!! cannot start server, err: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func MigrateDb() {
	if m, err := migrate.New(
		"file://"+viper.GetString("migrations_path"),
		"postgres://postgres:newpassword@0.0.0.0:5436/postgres?sslmode=disable"); err != nil {
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
