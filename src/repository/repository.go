package repository

import (
	"github.com/jphacks/A_2108/src/config"
	"github.com/jphacks/A_2108/src/database"
	"gorm.io/gorm"
)

const (
	CANNOT_CONVERT    = "Recived type can not be converted into MaskedUser"
	RECORD_NOT_FOUND  = "Record is not found"
	CREATOR_NOT_FOUND = "Creator is not found"
)

func errHandling(err error) error {
	if err == gorm.ErrRecordNotFound {
		return &PlanError{RECORD_NOT_FOUND}
	}
	return err
}

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
