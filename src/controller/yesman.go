package controller

import (
	"github.com/jphacks/A_2108/src/api_response"
	"github.com/jphacks/A_2108/src/domain"
)

type yesmanUserRepository struct{}
type yesmanPlanRepository struct{}

func (ur *yesmanUserRepository) GetUserByID(id int) (api_response.User, error) {
	return api_response.User{}, nil
}
func (ur *yesmanUserRepository) PostUser(user domain.User) (int, error) {
	return 0, nil
}
func (ur *yesmanUserRepository) PutUser(user domain.User) error {
	return nil
}
func (ur *yesmanUserRepository) DeleteUserByID(id int) error {
	return nil
}

func (pr *yesmanPlanRepository) GetPlansOrderedbyTime(limit int) (domain.Plans, error) {
	return nil, nil
}

func (pr *yesmanPlanRepository) GetPlanByID(id int) (domain.Plan, error) {
	return domain.Plan{}, nil
}

func (pr *yesmanPlanRepository) PostPlan(domain.Plan) (int, error) {
	return 0, nil
}

func (pr *yesmanPlanRepository) PutPlan(api_response.Plan) error {
	return nil
}

func (pr *yesmanPlanRepository) DeletePlanByID(id int) error {
	return nil
}
