package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/deadlorpa/auth-app/appconfig"
	"github.com/deadlorpa/auth-app/handler"
	"github.com/deadlorpa/auth-app/model"
	"github.com/deadlorpa/auth-app/repository"
	"github.com/deadlorpa/auth-app/service"
)

func main() {
	config, err := appconfig.Get()
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
	go func() {
		if err := srv.Run(config.Host, handlers.InitRoutes()); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("!!! cannot start server, err: %s", err.Error())
			}
		}
	}()

	log.Print("~~~ service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("~~~ service shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("!!! error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("!!! error occured on db connection close: %s", err.Error())
	}
}
