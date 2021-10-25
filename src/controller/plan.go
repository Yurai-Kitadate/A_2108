package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/domain"
)

func (con *Controller) PlanGet(c *gin.Context) {
	res := domain.Plans{
		Plans: []domain.Plan{
			{
				PlanId:      1,
				Title:       "title",
				Description: "description",
				Image:       "url",
				Creator: &domain.Creator{
					ID:          1,
					Image:       "url",
					DisplayName: "name",
					Job: &domain.Job{
						ID:             1,
						Jobname:        "job",
						Dateoffirstjob: time.Now(),
					},
				},
			},
		},
	}
	c.JSON(200, res)
}

func (con *Controller) PlanGetPathParam(c *gin.Context) {
	planId := c.Param("id")
	planIdInt, err := strconv.Atoi(planId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": "Atoi error: " + err.Error(),
		})
		return
	}
	res := domain.Plans{
		Plans: []domain.Plan{
			{
				PlanId:      planIdInt,
				Title:       "title",
				Description: "description",
				Image:       "url",
				Creator: &domain.Creator{
					ID:          1,
					Image:       "url",
					DisplayName: "name",
					Job: &domain.Job{
						ID:             1,
						Jobname:        "job",
						Dateoffirstjob: time.Now(),
					},
				},
				Days: domain.Days{
					{
						Headings: domain.Headings{
							{
								ID:    1,
								Text:  "text",
								Order: 1,
							},
						},
						Schedule: domain.Schedule{
							{
								ID:          1,
								Description: "text",
								StartTime:   time.Now(),
								EndTime:     time.Now(),
								Place: domain.Place{
									ID:         1,
									Area:       "area",
									Prefecture: "pref",
									City:       "city",
								},
								HpLink:          "link",
								ReservationLink: "link",
								Order:           1,
							},
						},
					},
				},
				Conditions: &domain.Conditions{
					ID: 1,
					Season: domain.Season{
						{
							ID:   1,
							Text: "text",
						},
					},
					TimeSpan: domain.TimeSpan{
						{
							ID:   1,
							Text: "text",
						},
					},
					Category: domain.Category{
						{
							ID:   1,
							Text: "text",
						},
					},
				},
			},
		},
	}
	c.JSON(200, res)
}

func (con *Controller) PlanPost(c *gin.Context) {
	var req domain.Plan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%#v\n", req)

	var res interface{}
	c.JSON(200, res)
}

func (con *Controller) PlanDelete(c *gin.Context) {
	planId := c.Param("id")
	fmt.Printf("PlanID: %v\n", planId)

	c.JSON(200, nil)
}
