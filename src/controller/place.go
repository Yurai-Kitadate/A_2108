package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/place"
)

func (con *Controller) GetPlace(c *gin.Context) {
	area, ok := c.GetQuery("area")
	if !ok {
		c.JSON(400, gin.H{"error": "area is required"})
		return
	}
	pref, ok := c.GetQuery("pref")
	if !ok {
		c.JSON(400, gin.H{"error": "pref is required"})
		return
	}

	intArea, err := strconv.Atoi(area)
	if err != nil {
		c.JSON(400, gin.H{"error": "area is invalid"})
		return
	}
	intPref, err := strconv.Atoi(pref)
	if err != nil {
		c.JSON(400, gin.H{"error": "pref is invalid"})
		return
	}

	place := place.GetPlace(intArea, intPref)
	c.JSON(200, place)
}
