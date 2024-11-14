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

	appController := internal.NewAppController(db)
	e := echo.New()
	e = internal.NewRouter(e, appController)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
