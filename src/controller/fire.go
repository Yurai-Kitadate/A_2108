package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/domain"
)

func (con *Controller) Fire1(c *gin.Context) {
	res := domain.Fire1{
		A: 1,
		B: "mieruka?",
		C: "mienaiyo",
	}
	c.JSON(200, res)
}
