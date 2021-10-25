package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type creator struct {
	ID          int    `json:"id,omitempty"`
	Image       string `json:"image,omitempty"`
	DisplayName string `json:"displayName"`
	Job         job    `json:"job,omitempty"`
}
type headings []struct {
	ID    int    `json:"id,omitempty"`
	Text  string `json:"text"`
	Order int    `json:"order"`
}
type place struct {
	ID         int    `json:"id,omitempty"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
}
type schedule []struct {
	ID              int       `json:"id,omitempty"`
	Description     string    `json:"description"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"EndTime"`
	Place           place     `json:"place"`
	HpLink          string    `json:"hpLink"`
	ReservationLink string    `json:"reservationLink"`
	Order           int       `json:"order"`
}
type days []struct {
	Headings headings `json:"headings"`
	Schedule schedule `json:"schedule"`
}
type season []struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}
type timeSpan []struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}
type category []struct {
	ID   int    `json:"id,omitempty"`
	Text string `json:"text"`
}
type conditions struct {
	ID       int      `json:"id,omitempty"`
	Season   season   `json:"season"`
	TimeSpan timeSpan `json:"timeSpan"`
	Category category `json:"category"`
}
type plan struct {
	PlanId      int         `json:"planId,omitempty"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Creator     *creator    `json:"creator,omitempty"`
	Days        days        `json:"days,omitempty"`
	Conditions  *conditions `json:"conditions,omitempty"`
}
type plans struct {
	Plans []plan `json:"plans"`
}

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
					Job: job{
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
					Job: job{
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
