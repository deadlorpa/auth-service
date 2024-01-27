package main

import (
	"flag"
	"log"

	"github.com/deadlorpa/auth-app/configs"
	"github.com/deadlorpa/auth-app/repository"
)

func main() {
	var method string
	flag.StringVar(&method, "m", "up", "Specify migration method. Default is up")
	flag.Parse()

	config, err := configs.Get()
	if err != nil {
		log.Fatalf("!!! cannot read configs: %s", err.Error())
	}

	repository.MigrateDB(config.DBConfig, method)
}
