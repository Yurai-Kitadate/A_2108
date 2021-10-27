package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/api_response"
	"github.com/jphacks/A_2108/src/domain"
)

type UserRepository interface {
	GetUserByID(string) (api_response.User, error)
	PostUser(api_response.User) (int, error)
	PutUser(api_response.User) error
	DeleteUserByID(int) error
}

func (con *Controller) UserGet(c *gin.Context) {
	res := domain.User{
		ID:          0,
		UserName:    "username",
		Image:       "url",
		Email:       "email",
		DisplayName: "displayName",
		Birthday:    time.Now(),
		Sex:         "sex",
		Contacts: domain.Contacts{
			ID:        0,
			Hp:        "hp",
			Instagram: "insta",
			Twitter:   "twitter",
			Facebook:  "facebook",
			Tiktok:    "tiktok",
			Biography: "bio",
		},
		Creator: domain.Creator{
			DisplayName: "displayName",
		},
		Job: domain.Job{
			ID:             0,
			Jobname:        "jobname",
			Dateoffirstjob: time.Now(),
		},
		Address: domain.Address{
			ID:          0,
			Area:        "会津",
			Prefecture:  "福島県",
			City:        "会津若松",
			Description: "Hey!",
		},
	}
	c.JSON(200, res)
}
