package controller

import "github.com/jphacks/A_2108/src/repository"

/* type Controller struct {
	UserRepository  UserRepository
	PlanRepository  PlanRepository
	ImageRepository ImageRepository
} */

type Controller struct {
	UserRepository  *repository.UserRepository
	PlanRepository  PlanRepository
	ImageRepository ImageRepository
}

func NewController() *Controller {
	return &Controller{}
}

/* func NewControllerWithYesmanRepository() *Controller {
	return &Controller{
		UserRepository:  &yesmanUserRepository{},
		PlanRepository:  &yesmanPlanRepository{},
		ImageRepository: &yesmanImageRepository{},
	}
} */

func NewControllerWithRepository() *Controller {
	return &Controller{
		UserRepository:  &repository.UserRepository{},
		PlanRepository:  &yesmanPlanRepository{},
		ImageRepository: &yesmanImageRepository{},
	}
}

func wrapToken(token string) map[string]string {
	return map[string]string{
		"token": token,
	}
}
