package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func intParam(c *gin.Context, param string) (int, error) {
	str := c.Param(param)
	return strconv.Atoi(str)
}
