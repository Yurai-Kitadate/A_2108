package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/domain"
)

type PlanRepository interface {
	GetPlansOrderedbyTime(int) (domain.Plans, error)
	GetPlanByID(int) (domain.Plan, error)
	PostPlan(domain.Plan) (int, error)
	PutPlan(domain.Plan) error
	DeletePlanByID(int) error
}

func (con *Controller) GetAllPlans(c *gin.Context) {
	plans, err := con.PlanRepository.GetPlansOrderedbyTime(-1)
	if err != nil {
		AbortWithError(c, http.StatusBadRequest, "Database Error", err)
		return
	}
	for i, plan := range plans {
		user, ok := plan.CreatorUser.(domain.User)
		if ok {
			plans[i].CreatorUser = user.Masked()
		}
	}
	c.JSON(200, plans)
}

func (con *Controller) GetPlanByID(c *gin.Context) {
	planID, err := intParam(c, "id")

	if err != nil {
		AbortWithError(c, http.StatusBadRequest, "Atoi Error", err)
		return
	}

	plan, err := con.PlanRepository.GetPlanByID(planID)

	if err != nil {
		AbortWithError(c, http.StatusBadRequest, "Database Error", err)
		return
	}
	{
		user, ok := plan.CreatorUser.(domain.User)
		if ok {
			plan.CreatorUser = user.Masked()
		}
	}

	c.JSON(200, plan)
}

func (con *Controller) CreatePlan(c *gin.Context) {
	var plan domain.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		AbortWithError(c, http.StatusBadRequest, "Bad Plan", err)
		return
	}

	id, err := con.PlanRepository.PostPlan(plan)

	if err != nil {
		AbortWithError(c, http.StatusInternalServerError, "Failed to add user", err)
		return
	}
	c.JSON(200, map[string]int{"id": id})
}

func (con *Controller) UpdatePlan(c *gin.Context) {
	var plan domain.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		AbortWithError(c, http.StatusBadRequest, "Bad Plan", err)
		return
	}

	err := con.PlanRepository.PutPlan(plan)

	if err != nil {
		AbortWithError(c, http.StatusInternalServerError, "Faield to add plan", err)
		return
	}
	c.JSON(200, map[string]string{"message": "Successful update plan"})
}

func (con *Controller) DeletePlanByID(c *gin.Context) {
	planID, err := intParam(c, "id")

	if err != nil {
		AbortWithError(c, http.StatusBadRequest, "Atoi Error", err)
		return
	}

	err = con.PlanRepository.DeletePlanByID(planID)

	if err != nil {
		AbortWithError(c, http.StatusBadRequest, "Database Error", err)
		return
	}

	c.JSON(200, map[string]string{"message": "Successful delete plan"})
}

func (con *Controller) PlanGet(c *gin.Context) {
	res := domain.Plans{
		{
			//			PlanId:      1,
			//			Title:       "title",
			//			Description: "description",
			//			Image:       "url",
			//			Creator: &domain.Creator{
			//				ID:          1,
			//				Image:       "url",
			//				DisplayName: "name",
			//				Job: &domain.Job{
			//					ID:             1,
			//					Jobname:        "job",
			//					Dateoffirstjob: time.Now(),
			//				},
			//			},
		},
	}
	c.JSON(200, res)
}

func (con *Controller) PlanGetPathParam(c *gin.Context) {
	planId := c.Param("id")
	_, err := strconv.Atoi(planId)
	if err != nil {
		AbortWithError(c, http.StatusBadRequest, "Atoi Error", err)
		return
	}
	res := domain.Plan{
		//PlanId:      planIdInt,
		//Title:       "title",
		//Description: "description",
		//Image:       "url",
		//Creator: &domain.Creator{
		//	ID:          1,
		//	Image:       "url",
		//	DisplayName: "name",
		//	Job: &domain.Job{
		//		ID:             1,
		//		Jobname:        "job",
		//		Dateoffirstjob: time.Now(),
		//	},
		//},
		//Days: domain.Days{
		//	{
		//		Headings: domain.Headings{
		//			{
		//				ID:    1,
		//				Text:  "text",
		//				Order: 1,
		//			},
		//		},
		//		Schedule: domain.Schedule{
		//			{
		//				ID:          1,
		//				Description: "text",
		//				StartTime:   time.Now(),
		//				EndTime:     time.Now(),
		//				Place: domain.Place{
		//					ID:         1,
		//					Area:       "area",
		//					Prefecture: "pref",
		//					City:       "city",
		//				},
		//				HpLink:          "link",
		//				ReservationLink: "link",
		//				Order:           1,
		//			},
		//		},
		//	},
		//},
		//Conditions: &domain.Conditions{
		//	ID: 1,
		//	Season: domain.Season{
		//		{
		//			ID:   1,
		//			Text: "text",
		//		},
		//	},
		//	TimeSpan: domain.TimeSpan{
		//		{
		//			ID:   1,
		//			Text: "text",
		//		},
		//	},
		//	Category: domain.Category{
		//		{
		//			ID:   1,
		//			Text: "text",
		//		},
		//	},
		//},
	}
	c.JSON(200, res)
}

func (con *Controller) PlanPost(c *gin.Context) {
	var req domain.Plan
	if err := c.ShouldBindJSON(&req); err != nil {
		AbortWithError(c, http.StatusBadRequest, "Bad Plan", err)
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
