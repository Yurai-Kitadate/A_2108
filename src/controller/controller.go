package controller

import (
	"github.com/gin-gonic/gin"
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
		UserRepository:  repository.NewUserRepository(db),
		PlanRepository:  repository.NewPlanRepository(db),
		ImageRepository: &yesmanImageRepository{}, // TODO
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

func AbortWithError(c *gin.Context, statusCode int, obj interface{}, e error) {
	var res map[string]interface{} = make(map[string]interface{})
	res["error"] = obj
	if gin.Mode() != "release" {
		res["errorData"] = e
	}
	c.JSON(statusCode, res)
}
