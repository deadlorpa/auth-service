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
	if err := InitConfig(); err != nil {
		log.Fatalf("!!! cannot read configs: %s", err.Error())
	}

	var dbCofig = configs.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	}

	repository.MigrateDB(dbCofig)

	db, err := repository.NewPostgresDB(dbCofig)
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

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
