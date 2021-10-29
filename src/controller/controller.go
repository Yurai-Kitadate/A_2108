package controller

import (
	"github.com/jphacks/A_2108/src/repository"
	"gorm.io/gorm"
)

type Controller struct {
	UserRepository  UserRepository
	PlanRepository  PlanRepository
	ImageRepository ImageRepository
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		UserRepository: repository.NewUserRepository(db),
		PlanRepository: repository.NewPlanRepository(db),
		// ImageRepositoryの代入をしていないので，Image回りは動かないです
	}
}

func NewControllerWithYesmanRepository() *Controller {
	return &Controller{
		UserRepository:  &yesmanUserRepository{},
		PlanRepository:  &yesmanPlanRepository{},
		ImageRepository: &yesmanImageRepository{},
	}
}

func wrapToken(token string) map[string]string {
	return map[string]string{
		"token": token,
	}
}
