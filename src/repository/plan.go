package repository

import (
	"github.com/jphacks/A_2108/src/api_response"
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

func (pr PlanRepository) GetPlanByID(planID int) (domain.Plan, error) {
	db := pr.db
	plan := domain.Plan{}

	{
		db_plan := domain.DBPlan{}

		err := db.First(&db_plan).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		plan.PlanId = db_plan.ID
		plan.Title = db_plan.Title
		plan.Description = db_plan.Description
		plan.Image = db_plan.Image
		plan.CreatedAt = db_plan.CreatedAt
	}

	{
		db_days := []domain.DBDay{}

		err := db.Find(&db_days).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}

		plan.Days = make(domain.Days, len(db_days))
		for i, v := range db_days {
			plan.Days[i].
				plan.Days[i].NthDay = v.NthDay
		}
	}

	{
		db_heading := []domain.DBHeading{}

		err := db.Where("day_id = ?")
	}

	// TODO Category Definitonをちゃんと設定する.
	return plan, nil
}

func (pr PlanRepository) GetPlansOrderedbyTime(limit int) ([]domain.DBPlan, error) {
	return nil, nil
}

func (pr PlanRepository) DeletePlanByID(planID int) error {
	return nil
}

func (pr PlanRepository) PostPlan(plan api_response.Plan) (int, error) {
	return -1, nil
}

func (pr PlanRepository) PutPlan(plan api_response.Plan) error {
	return nil
}
