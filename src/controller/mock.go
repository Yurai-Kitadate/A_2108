package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/domain"
	"github.com/jphacks/A-2108/src/mock"
)

func MockGetUserByID(c *gin.Context) {
	planId := c.Param("id")
	planIdInt, err := strconv.Atoi(planId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Atoi error: " + err.Error(),
		})
		return
	}

	var res domain.User
	if planIdInt == 1 {
		res = mock.MockUser1
	} else if planIdInt == 2 {
		res = mock.MockUser2
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Not Found",
		})
	}
	c.JSON(200, res)
}

func MockGetPlanByID(c *gin.Context) {
	planId := c.Param("id")
	planIdInt, err := strconv.Atoi(planId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Atoi error: " + err.Error(),
		})
		return
	}

	for _, plan := range domain.MockPlans {
		if plan.PlanId == planIdInt {
			c.JSON(200, plan)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"Error": "Not Found",
	})
}

func MockGetAllPlans(c *gin.Context) {
	c.JSON(200, mock.MockPlans)
}
