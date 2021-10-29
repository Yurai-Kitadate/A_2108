package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/place"
)

func (con *Controller) GetPlace(c *gin.Context) {
	intArea := 0
	intPref := 0

	area, ok := c.GetQuery("area")
	var err error
	if ok {
		intArea, err = strconv.Atoi(area)
		if err != nil {
			c.JSON(400, gin.H{"error": "area is invalid"})
			return
		}
	}
	pref, ok := c.GetQuery("pref")
	if ok {
		intPref, err = strconv.Atoi(pref)
		if err != nil {
			c.JSON(400, gin.H{"error": "pref is invalid"})
			return
		}
	}

	place := place.GetPlace(intArea, intPref)
	c.JSON(200, place)
}
