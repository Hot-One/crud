package main

import (
	"fmt"
	"net/http"

	"user/api"
	"user/config"
	"user/storage/postgres"
)

func main() {

	cfg := config.Load()

	strg, err := postgres.NewConnectionPostgres(cfg)
	if err != nil {
		panic("No connection with database" + err.Error())
	}

	defer strg.Close()

	api.NewApi(&cfg, strg)

	fmt.Println("Listening..." + cfg.HTTPPort)
	err = http.ListenAndServe(cfg.ServerHost+cfg.HTTPPort, nil)
	if err != nil {
		panic("Panica qlish kere" + err.Error())
	}

}
