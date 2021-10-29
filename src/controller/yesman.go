package controller

import (
	"errors"

	"github.com/jphacks/A_2108/src/domain"
)

/* type yesmanUserRepository struct{} */
type yesmanPlanRepository struct{}
type yesmanImageRepository struct{}

/* func (ur *yesmanUserRepository) GetUserByID(id int) (domain.User, error) {
	if id == 1 {
		return MockUser1, nil
	}
	if id == 2 {
		return MockUser2, nil
	}
	return domain.User{}, errors.New("Not found")
}
func (ur *yesmanUserRepository) PostUser(user domain.User) (int, error) {
	return 0, nil
}
func (ur *yesmanUserRepository) PutUser(user domain.User) error {
	return nil
}
func (ur *yesmanUserRepository) DeleteUserByID(id int) error {
	return nil
} */

func (pr *yesmanPlanRepository) GetPlansOrderedbyTime(limit int) (domain.Plans, error) {
	return MockPlans, nil
}

func (pr *yesmanPlanRepository) GetPlanByID(id int) (domain.Plan, error) {
	for _, plan := range MockPlans {
		if plan.PlanId == id {
			return plan, nil
		}
	}
	return domain.Plan{}, errors.New("Not found")
}

func (pr *yesmanPlanRepository) PostPlan(domain.Plan) (int, error) {
	return 0, nil
}

func (pr *yesmanPlanRepository) PutPlan(domain.Plan) error {
	return nil
}

func (pr *yesmanPlanRepository) DeletePlanByID(id int) error {
	return nil
}

/* func (pr *yesmanUserRepository) GetUserByEmail(string) (domain.User, error) {
	return domain.User{}, nil
} */

func (ir *yesmanImageRepository) CreateUser(img domain.Image) error {
	return nil
}

func (ir *yesmanImageRepository) GetImagesByUserID(int) ([]domain.Image, error) {
	return nil, nil
}
