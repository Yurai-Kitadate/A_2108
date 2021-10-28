package controller

type Controller struct {
	UserRepository UserRepository
	PlanRepository PlanRepository
}

func NewController() *Controller {
	return &Controller{}
}

func NewControllerWithYesmanRepository() *Controller {
	return &Controller{
		UserRepository: &yesmanUserRepository{},
		PlanRepository: &yesmanPlanRepository{},
	}
}

func wrapToken(token string) map[string]string {
	return map[string]string{
		"token": token,
	}
}
