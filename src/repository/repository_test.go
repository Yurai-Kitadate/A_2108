package repository

import (
	"testing"

	"github.com/jphacks/A_2108/src/database"
	"github.com/jphacks/A_2108/src/mock"
)

func TestGetUserByID(t *testing.T) {
	db, err := database.NewDatabaseHandlerWithDBName("DAWN")
	if err != nil {
		t.Errorf("DB Open Error: %+v", err)
	}

	err = DrivePostUser(db, mock.MockUser1)
	if err != nil {
		t.Errorf("%+v", err)
	}

	err = DriveGetUserByID(db, 1)
	if err != nil && err.Error() != "Not Creator" {
		t.Errorf("%+v", err)
	}

	err = DrivePostPlan(db, mock.MockPlan)
	if err != nil {
		t.Errorf("%+v", err)
	}

	err = DriveGetPlanbyID(db, 1)
	if err != nil {
		t.Errorf("%+v", err)
	}
}
