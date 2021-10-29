package repository

import (
	"fmt"

	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

func DriveAutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.DBUser{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBCreator{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBJob{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBContacts{}); err != nil {
		panic("Migration Error")
	}

	if err := db.AutoMigrate(&domain.DBPlan{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBCondition{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBSeason{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBTimeSpan{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBCategory{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBSeasonDefinition{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBTimeSpanDefinition{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBCategoryDefinition{}); err != nil {
		panic("Migration Error")
	}

	if err := db.AutoMigrate(&domain.DBDay{}); err != nil {
		panic("Migration Error")
	}

	if err := db.AutoMigrate(&domain.DBHeading{}); err != nil {
		panic("Migration Error")
	}

	if err := db.AutoMigrate(&domain.DBSchedule{}); err != nil {
		panic("Migration Error")
	}

	if err := db.AutoMigrate(&domain.DBPlace{}); err != nil {
		panic("Migration Error")
	}
	if err := db.AutoMigrate(&domain.DBAddress{}); err != nil {
		panic("Migration Error")
	}
}

func DriveInitializeDefinitions(db *gorm.DB) {
	seasonDef := []domain.DBSeasonDefinition{
		{Description: "1月"},
		{Description: "2月"},
		{Description: "3月"},
		{Description: "4月"},
		{Description: "5月"},
		{Description: "6月"},
		{Description: "7月"},
		{Description: "8月"},
		{Description: "9月"},
		{Description: "10月"},
		{Description: "11月"},
		{Description: "12月"},
	}
	categoryDef := []domain.DBCategoryDefinition{
		{Description: "れじゃー"},
		{Description: "すぽーつ"},
		{Description: "おんせん"},
		{Description: "れきしめぐり"},
		{Description: "たべあるき"},
	}
	timespanDef := []domain.DBTimeSpanDefinition{
		{Description: "1泊"},
		{Description: "2泊"},
		{Description: "3泊"},
		{Description: "4泊"},
		{Description: "5泊"},
		{Description: "6泊"},
		{Description: "7泊"},
		{Description: "8泊"},
		{Description: "9泊"},
		{Description: "10泊いじょー"},
	}

	db.Create(&seasonDef)
	db.Create(&categoryDef)
	db.Create(&timespanDef)
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

func DrivePostUser(db *gorm.DB, user domain.User) error {
	user_repository := UserRepository{db}
	res, err := user_repository.PostUser(user)

	if err != nil {
		return err
	}
	fmt.Printf("Added UserID: %d", res)
	return nil
}

func DrivePostPlan(db *gorm.DB, plan domain.Plan) error {
	pr := PlanRepository{db}
	res, err := pr.PostPlan(plan)

	if err != nil {
		return err
	}
	fmt.Printf("Added UserID: %d", res)
	return nil
}

func DriveGetPlanbyID(db *gorm.DB, planID int) error {
	pr := PlanRepository{db}
	res, err := pr.GetPlanByID(planID)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", res)
	return nil
}
