package controller

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/auth"
	"github.com/jphacks/A_2108/src/domain"
	"github.com/jphacks/A_2108/src/s3"
)

type ImageRepository interface {
	CreateUser(domain.Image) error
	GetImagesByUserID(int) ([]domain.Image, error)
}

func (con *Controller) ImagePost(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: IDの取得を実装
	id, err := auth.GetIdBySession(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sess := s3.NewS3Session()
	endpoint, err := s3.UploadToS3(sess, body, strconv.Itoa(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, endpoint)
}

func (con *Controller) GetOwnImages(c *gin.Context) {
	id, err := auth.GetIdBySession(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	images, err := con.ImageRepository.GetImagesByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, images)
}
