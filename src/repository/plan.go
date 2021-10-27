package repository

import (
	"github.com/jphacks/A_2108/src/api_response"
	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type planRepository struct {
	db *gorm.DB
}

func (plan_repository planRepository) GetPlanByID(planID int) (domain.DBPlan, error) {
	return domain.DBPlan{}, nil
}

func (plan_repository planRepository) GetPlanOrderedbyTime(limit int) ([]domain.DBPlan, error) {
	return nil, nil
}

func (plan_repository planRepository) DeletePlanByID(planID int) error {
	return nil
}

func (plan_repository planRepository) PostPlan(plan api_response.Plan) (int, error) {
	return -1, nil
}

func (plan_repository planRepository) PutPlan(plan api_response.Plan) error {
	return nil
}
