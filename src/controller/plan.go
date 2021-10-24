package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type creator struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	DisplayName string `json:"displayName"`
	Job         job    `json:"job"`
}
type headings []struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Order int    `json:"order"`
}
type place struct {
	ID         int    `json:"id"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
}
type schedule []struct {
	ID              int       `json:"id"`
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
	ID   int    `json:"id"`
	Text string `json:"text"`
}
type timeSpan []struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
type category []struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
type conditions struct {
	ID       int      `json:"id"`
	Season   season   `json:"season"`
	TimeSpan timeSpan `json:"timeSpan"`
	Category category `json:"category"`
}
type plan struct {
	PlanId      int         `json:"planId"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Creator     creator     `json:"creator"`
	Days        days        `json:"days,omitempty"`
	Conditions  *conditions `json:"conditions,omitempty"`
}
type plans struct {
	Plans []plan `json:"plans"`
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

func (con *Controller) PlanGetPathParam(c *gin.Context) {
	planId := c.Param("id")
	planIdInt, err := strconv.Atoi(planId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": "Atoi error: " + err.Error(),
		})
	}
	res := plans{
		Plans: []plan{
			{
				PlanId:      planIdInt,
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
				Days: days{
					{
						Headings: headings{
							{
								ID:    0,
								Text:  "text",
								Order: 0,
							},
						},
						Schedule: schedule{
							{
								ID:          0,
								Description: "text",
								StartTime:   time.Now(),
								EndTime:     time.Now(),
								Place: place{
									ID:         0,
									Area:       "area",
									Prefecture: "pref",
									City:       "city",
								},
								HpLink:          "link",
								ReservationLink: "link",
								Order:           0,
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
