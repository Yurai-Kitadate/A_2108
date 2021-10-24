package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

type creator struct {
	ID          int
	Image       string
	DisplayName string
	Job         job
}
type plan struct {
	PlanId      int
	Title       string
	Description string
	Image       string
	Creator     creator
}

func (con *Controller) PlanGet(c *gin.Context) {
	res := []plan{{
		PlanId:      0,
		Title:       "title",
		Description: "description",
		Image:       "url",
		Creator: creator{
			ID:          0,
			Image:       "url",
			DisplayName: "name",
			Job: job{
				ID:             0,
				Jobname:        "job",
				Dateoffirstjob: time.Now(),
			},
		},
	}}
	c.JSON(200, res)
}
