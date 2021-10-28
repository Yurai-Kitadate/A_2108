package repository

import (
	"fmt"
	"testing"

	"github.com/jphacks/A_2108/src/database"
	"github.com/jphacks/A_2108/src/domain"
)

func TestGetUserByID(t *testing.T) {
	db, err := database.NewDatabaseHandlerWithDBName("DAWN")
	if err != nil {
		fmt.Errorf("DB Open Error: %+v", err)
	}

	db.AutoMigrate(&domain.DBUser{})
	db.AutoMigrate(&domain.DBCreator{})
	db.AutoMigrate(&domain.DBJob{})
	db.AutoMigrate(&domain.DBContacts{})
	db.AutoMigrate(&domain.DBCondition{})
	db.AutoMigrate(&domain.DBSeasonDefinition{})
	db.AutoMigrate(&domain.DBTimeSpanDefinition{})
	db.AutoMigrate(&domain.DBCategoryDefinition{})
	db.AutoMigrate(&domain.DBDay{})
	db.AutoMigrate(&domain.DBHeading{})
	db.AutoMigrate(&domain.DBSchedule{})
	db.AutoMigrate(&domain.DBPlace{})

	err = DriveGetUserByID(db, 1)
	if err != nil {
		fmt.Errorf("%+v", err)
	}

	err = DrivePostUser(db)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
}
