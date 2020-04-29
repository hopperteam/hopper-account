package main

import (
	"github.com/hopperteam/hopper-account/config"
	"github.com/hopperteam/hopper-account/model"
	"github.com/hopperteam/hopper-account/security"
	"github.com/hopperteam/hopper-account/web"
	log "github.com/sirupsen/logrus"
	"os"
)



func main() {
	log.SetOutput(os.Stdout)
	log.Info("Initializing hopper-account")

	err := security.LoadKeys()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = model.ConnectDB(config.Config.DbConnectionStr, config.Config.DbName)
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := web.NewServer()
	log.Info("Starting web server")
	srv.Start()
}
