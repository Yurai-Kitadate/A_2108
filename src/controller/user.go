package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

type contacts struct {
	ID        int    `json:"id"`
	Hp        string `json:"hp"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Tiktok    string `json:"tiktok"`
	Biography string `json:"biography"`
}
type job struct {
	ID             int       `json:"id"`
	Jobname        string    `json:"jobname"`
	Dateoffirstjob time.Time `json:"dateoffirstjob"`
}
type address struct {
	ID          int    `json:"id"`
	Area        string `json:"area"`
	Prefecture  string `json:"prefecture"`
	City        string `json:"city"`
	Description string `json:"description"`
}
type user struct {
	ID          int       `json:"id"`
	UserName    string    `json:"userName"`
	Image       string    `json:"image"`
	Email       string    `json:"email"`
	DisplayName string    `json:"displayName"`
	Birthday    time.Time `json:"birthday"`
	Sex         string    `json:"sex"`
	Contacts    contacts  `json:"contacts"`
	Creator     creator   `json:"creator"`
	Job         job       `json:"job"`
	Address     address   `json:"address"`
}

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
