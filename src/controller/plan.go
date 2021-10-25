package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (con *Controller) PlanGet(c *gin.Context) {
	res := plans{
		Plans: []plan{
			{
				PlanId:      1,
				Title:       "title",
				Description: "description",
				Image:       "url",
				Creator: &creator{
					ID:          1,
					Image:       "url",
					DisplayName: "name",
					Job: &job{
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
	res := plans{
		Plans: []plan{
			{
				PlanId:      planIdInt,
				Title:       "title",
				Description: "description",
				Image:       "url",
				Creator: &creator{
					ID:          1,
					Image:       "url",
					DisplayName: "name",
					Job: &job{
						ID:             1,
						Jobname:        "job",
						Dateoffirstjob: time.Now(),
					},
				},
				Days: days{
					{
						Headings: headings{
							{
								ID:    1,
								Text:  "text",
								Order: 1,
							},
						},
						Schedule: schedule{
							{
								ID:          1,
								Description: "text",
								StartTime:   time.Now(),
								EndTime:     time.Now(),
								Place: place{
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
				Conditions: &conditions{
					ID: 1,
					Season: season{
						{
							ID:   1,
							Text: "text",
						},
					},
					TimeSpan: timeSpan{
						{
							ID:   1,
							Text: "text",
						},
					},
					Category: category{
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
	var req plan
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
