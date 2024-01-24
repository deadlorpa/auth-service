package main

import (
	"log"
	"os"

	"github.com/deadlorpa/auth-app/handler"
	"github.com/deadlorpa/auth-app/repository"
	"github.com/deadlorpa/auth-app/service"
	"github.com/deadlorpa/auth-app/structure"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("!!! cannot read configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("!!! error loading env variables: %s", err.Error())
	}

	var cfg = repository.Config{
		Host:           viper.GetString("db.host"),
		Port:           viper.GetString("db.port"),
		Username:       viper.GetString("db.username"),
		DBName:         viper.GetString("db.dbname"),
		SSLMode:        viper.GetString("db.sslmode"),
		Password:       os.Getenv("DB_PASSWORD"),
		MigrationsPath: viper.GetString("migrations_path"),
	}

	repository.MigrateDB(cfg)

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("!!! failed to initialize db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
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
