package internal

import (
	"parcial2-ingweb/internal/database"
	"parcial2-ingweb/internal/entity1"
	"parcial2-ingweb/internal/entity2"
)

type AppController struct {
	Entity1 interface{ entity1.IE1Controller }
	Entity2 interface{ entity2.IE2Controller }
}

func NewAppController(db *database.Connection) AppController {
	return AppController{
		Entity1: entity1.NewE1Controller(entity1.NewE1Interactor(entity1.NewE1Repository(db), entity1.NewE1Presenter())),
		Entity2: entity2.NewE2Controller(entity2.NewE2Interactor(entity2.NewE2Repository(db), entity2.NewE2Presenter())),
	}
}
