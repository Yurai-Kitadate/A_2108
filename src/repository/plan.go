package repository

import (
	"github.com/jphacks/A_2108/src/api_response"
	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type PlanRepository struct {
	db *gorm.DB
}

func (plan_repository PlanRepository) GetPlanByID(planID int) (domain.DBPlan, error) {
	db := plan_repository.db
	plan := domain.DBPlan{}

	err := db.Preload("Condition").Preload("Season").
		Preload("TimeSpan").Preload("Category").Preload("Day").
		Preload("Heading").Preload("Schedule").Preload("Address").
		First(&plan).Error

	if err == gorm.ErrRecordNotFound {
		return domain.DBPlan{}, &NotFoundError{}
	}

	// TODO Category Definitonをちゃんと設定する.
	return plan, nil
}

func (plan_repository PlanRepository) GetPlansOrderedbyTime(limit int) ([]domain.DBPlan, error) {
	db := plan_repository.db
	return nil, nil
}

func (plan_repository PlanRepository) DeletePlanByID(planID int) error {
	db := plan_repository.db
	return nil
}

func (plan_repository PlanRepository) PostPlan(plan api_response.Plan) (int, error) {
	db := plan_repository.db
	return -1, nil
}

func (plan_repository PlanRepository) PutPlan(plan api_response.Plan) error {
	db := plan_repository.db
	return nil
}
