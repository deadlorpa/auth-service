package main

import (
	"log"

	"github.com/deadlorpa/auth-app/handler"
	"github.com/deadlorpa/auth-app/repository"
	"github.com/deadlorpa/auth-app/service"
	"github.com/deadlorpa/auth-app/structure"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("!!! cannot read configs: %s", err.Error())
	}
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
