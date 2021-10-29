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

	hashPass, err := auth.CreateHash(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Create hash error": err.Error()})
		return
	}
	req.Password = hashPass

	id, err := con.UserRepository.PostUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, wrapToken(token))
}
