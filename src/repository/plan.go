package repository

import (
	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type PlanRepository struct {
	db *gorm.DB
}

type PlanError struct {
	s string
}

func (e PlanError) Error() string {
	return e.s
}

func (plan_repository PlanRepository) GetPlanByID(planID int) (domain.DBPlan, error) {
	db := plan_repository.db
	plan := domain.DBPlan{}

	err := db.Preload("Condition").Preload("Season").
		Preload("TimeSpan").Preload("Category").Preload("Day").
		Preload("Heading").Preload("Schedule").Preload("Address").
		First(&plan).Error

	if err == gorm.ErrRecordNotFound {
		return domain.DBPlan{}, &PlanError{"Record Not Found"}
	}

	// TODO Category Definitonをちゃんと設定する.
	return plan, nil
}

func (plan_repository PlanRepository) GetPlansOrderedbyTime(limit int) ([]domain.DBPlan, error) {
	return nil, nil
}

func (plan_repository PlanRepository) DeletePlanByID(planID int) error {
	return nil
}

func (plan_repository PlanRepository) PostPlan(plan domain.Plan) (int, error) {
	return -1, nil
}

func (plan_repository PlanRepository) PutPlan(plan domain.Plan) error {
	return nil
}
