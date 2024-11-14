package main

import (
	"log"
	"parcial2-ingweb/config"
	"parcial2-ingweb/internal"
	"parcial2-ingweb/internal/database"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.ProjectID)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}
	defer db.Close()

	controller := internal.NewController(internal.NewInteractor(internal.NewRepository(db), internal.NewPresenter()))
	e := echo.New()
	e = internal.NewRouter(e, controller)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
