package repository

import (
	"github.com/jphacks/A_2108/src/config"
	"github.com/jphacks/A_2108/src/database"
)

var (
	once bool
)

func init() {
	db, err := database.NewDatabaseHandlerWithDBName("DAWN")
	if err != nil {
		panic("Can not connect DB.")
	}

	if config.IsTest() || config.IsTestonDocker() || !once {
		DriveAutoMigrate(db)

		DriveInitializeDefinitions(db)

		initPlan(db)
		once = true
	}
}
