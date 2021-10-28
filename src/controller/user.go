package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/api_response"
	"github.com/jphacks/A_2108/src/domain"
)

type UserRepository interface {
	GetUserByID(int) (api_response.User, error)
	PostUser(domain.User) (int, error)
	PutUser(api_response.User) error
	DeleteUserByID(int) error
}

func (con *Controller) GetUserByID(c *gin.Context) {
	userID, err := intParam(c, "id")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Atoi error: " + err.Error(),
		})
		return
	}

	user, err := con.UserRepository.GetUserByID(userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Storage Server Error: ",
		})
	}
	c.JSON(200, user)
}

func (con *Controller) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad user"})
		return
	}
	id, err := con.UserRepository.PostUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}
	c.JSON(200, map[string]int{"id": id})
}

func (con *Controller) UserGet(c *gin.Context) {
	res := domain.User{
		//		ID:          0,
		//		UserName:    "username",
		//		Image:       "url",
		//		Email:       "email",
		//		DisplayName: "displayName",
		//		Birthday:    time.Now(),
		//		Sex:         "sex",
		//		Contacts: domain.Contacts{
		//			ID:        0,
		//			Hp:        "hp",
		//			Instagram: "insta",
		//			Twitter:   "twitter",
		//			Facebook:  "facebook",
		//			Tiktok:    "tiktok",
		//			Biography: "bio",
		//		},
		//		Creator: domain.Creator{
		//			DisplayName: "displayName",
		//		},
		//		Job: domain.Job{
		//			ID:             0,
		//			Jobname:        "jobname",
		//			Dateoffirstjob: time.Now(),
		//		},
		//		Address: domain.Address{
		//			ID:          0,
		//			Area:        "会津",
		//			Prefecture:  "福島県",
		//			City:        "会津若松",
		//			Description: "Hey!",
		//		},
	}
	c.JSON(200, res)
}
