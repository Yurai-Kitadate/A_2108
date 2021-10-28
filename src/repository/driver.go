package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

func DriveAutoMigrate(db *gorm.DB, t *testing.T) {
	err := db.AutoMigrate(&domain.DBUser{})
	if err != nil {
		t.Errorf("Migration Error")
	}
	err = db.AutoMigrate(&domain.DBCreator{})
	if err != nil {
		t.Errorf("Migration Error")
	}
	err = db.AutoMigrate(&domain.DBJob{})
	if err != nil {
		t.Errorf("Migration Error")
	}
	err = db.AutoMigrate(&domain.DBContacts{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBCondition{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBSeasonDefinition{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBTimeSpanDefinition{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBCategoryDefinition{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBDay{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBHeading{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBSchedule{})
	if err != nil {
		t.Errorf("Migration Error")
	}

	err = db.AutoMigrate(&domain.DBPlace{})
	if err != nil {
		t.Errorf("Migration Error")
	}
}

func DriveGetUserByID(db *gorm.DB, userID int) error {
	user_repository := UserRepository{db}
	res, err := user_repository.GetUserByID(userID)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", res)
	return nil
}

func DrivePostUser(db *gorm.DB) error {
	user_repository := UserRepository{db}
	user := domain.User{
		UserName:    "pachimon_Lightbells",
		Email:       "takahasi0928@gmail.com",
		Password:    "Password",
		Image:       "https://google.com",
		DisplayName: "らいとべる",
		DateOfBirth: time.Now(),
		Sex:         domain.Male,
		Contacts:    domain.Contacts{},
		Creator:     nil,
		Place:       domain.Place{},
	}
	res, err := user_repository.PostUser(user)

	if err != nil {
		return err
	}
	fmt.Printf("Added UserID: %d", res)
	return nil
}
