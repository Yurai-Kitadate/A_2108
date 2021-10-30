package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/auth"
	"github.com/jphacks/A_2108/src/domain"
)

type UserRepository interface {
	GetUserByID(int) (domain.User, error)
	PostUser(domain.User) (int, error)
	PutUser(domain.User) error
	DeleteUserByUserID(int) error
	GetUserByEmail(string) (domain.User, error)
	GetIsUniqueEmail(string) (bool, error)
	GetIsUniqueUserName(string) (bool, error)
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
	c.JSON(200, user.Masked())
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

func (con *Controller) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad user"})
		return
	}
	err := con.UserRepository.PutUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faield to update user"})
		return
	}
	c.JSON(200, map[string]string{"mesasge": "OK"})
}

func (con *Controller) DeleteUser(c *gin.Context) {
	userID, err := intParam(c, "id")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Atoi error: " + err.Error(),
		})
		return
	}

	err = con.UserRepository.DeleteUserByUserID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed delete user",
		})
		return
	}
	c.JSON(200, map[string]string{"message": "Successful delete user"})
}

func (con *Controller) IsValidEmail(c *gin.Context) {
	email := c.Param("email")
	ok, err := con.UserRepository.GetIsUniqueEmail(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": "Database Error",
		})
		return
	}
	c.JSON(200, map[string]bool{"ok": ok})
}

func (con *Controller) IsValidUserName(c *gin.Context) {
	email := c.Param("username")
	ok, err := con.UserRepository.GetIsUniqueUserName(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": "Database Error",
		})
		return
	}
	c.JSON(200, map[string]bool{"ok": ok})
}


func (con *Controller) DeleteCreator(c *gin.Context) {
	userID, err := auth.GetIdBySession(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Authorization error",
		})
	}

	err = con.UserRepository.DeleteUserByUserID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": "Faieled to delete creator",
		})
	}
	c.JSON(200, map[string]string{"message": "successful delete creator"})
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
