package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (con *Controller) UserGet(c *gin.Context) {
	res := user{
		ID:          0,
		UserName:    "username",
		Image:       "url",
		Email:       "email",
		DisplayName: "displayName",
		Birthday:    time.Now(),
		Sex:         "sex",
		Contacts: contacts{
			ID:        0,
			Hp:        "hp",
			Instagram: "insta",
			Twitter:   "twitter",
			Facebook:  "facebook",
			Tiktok:    "tiktok",
			Biography: "bio",
		},
		Creator: creator{
			DisplayName: "displayName",
		},
		Job: job{
			ID:             0,
			Jobname:        "jobname",
			Dateoffirstjob: time.Now(),
		},
		Address: address{
			ID:          0,
			Area:        "会津",
			Prefecture:  "福島県",
			City:        "会津若松",
			Description: "Hey!",
		},
	}
	c.JSON(200, res)
}
