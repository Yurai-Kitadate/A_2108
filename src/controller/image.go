package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (con *Controller) ImagePost(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%v\n", body)
	c.JSON(200, nil)
}
