package main

import (
	"log"

	"github.com/deadlorpa/auth-app/configs"
	"github.com/deadlorpa/auth-app/handler"
	"github.com/deadlorpa/auth-app/model"
	"github.com/deadlorpa/auth-app/repository"
	"github.com/deadlorpa/auth-app/service"

	"github.com/spf13/viper"
)

func main() {
	config, err := configs.Get()
	if err != nil {
		log.Fatalf("!!! cannot read configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config.DBConfig)
	if err != nil {
		log.Fatalf("!!! failed to initialize db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(model.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("!!! cannot start server, err: %s", err.Error())
	}
}
