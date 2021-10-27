package controller

import "github.com/jphacks/A_2108/src/api_response"

type yesmanUserRepository struct{}
type yesmanPlanRepository struct{}

func (ur *yesmanUserRepository) GetUserByID(id int) (api_response.User, error) {
	return api_response.User{}, nil
}
func (ur *yesmanUserRepository) PostUser(user api_response.User) (int, error) {
	return 0, nil
}
func (ur *yesmanUserRepository) PutUser(user api_response.User) error {
	return nil
}
func (ur *yesmanUserRepository) DeleteUserByID(id int) error {
	return nil
}

func (pr *yesmanPlanRepository) GetPlansOrderedbyTime(limit int) (api_response.Plans, error) {
	return nil, nil
}

func (pr *yesmanPlanRepository) GetPlanByID(id int) (api_response.Plan, error) {
	return api_response.Plan{}, nil
}

func (pr *yesmanPlanRepository) PostPlan(api_response.Plan) (int, error) {
	return 0, nil
}

func (pr *yesmanPlanRepository) PutPlan(api_response.Plan) error {
	return nil
}

func (pr *yesmanPlanRepository) DeletePlanByID(id int) error {
	return nil
}
