package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/auth"
	"github.com/jphacks/A_2108/src/domain"
)

func (con *Controller) RegisterPost(c *gin.Context) {
	var req domain.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: ユーザー追加処理
	/* id, err := PostUser(req)(int, error)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} */
	id := 101
	token, err := auth.GenerateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, token)
}
