package main

import (
	"app/api"
	"app/config"
	"app/pkg/db"
	"app/storage"
)

func main() {
	cfg := config.Load()

	pgxConn, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		return
	}

	storage := storage.NewStorage(pgxConn)

	api.Api(storage)
}
