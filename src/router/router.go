package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/auth"
	"github.com/jphacks/A_2108/src/controller"
)

func Route() *gin.Engine {
	r := gin.Default()
	con := controller.NewController()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/fire", con.Fire1)
	r.GET("/fire/pathpara/:id", con.FirePath)

	// 認証が必要なAPI
	authGroup := r.Group("/", auth.VerifyAPIMiddleware())
	{
		authGroup.GET("/user", con.UserGet)
		authGroup.GET("/user/:id", controller.MockGetUserByID)
		authGroup.GET("/plan", controller.MockGetAllPlans)
		authGroup.GET("/plan/:id", controller.MockGetPlanByID)
		//r.GET("/plan", con.PlanGet)
		//r.GET("/plan/:id", con.PlanGetPathParam)
		authGroup.POST("/plan", con.PlanPost)
		authGroup.DELETE("/plan/:id", con.PlanDelete)
		authGroup.POST("/image", con.ImagePost)
	}

	r.POST("/register", con.RegisterPost)
	r.POST("/login", con.LoginPOST)
	return r
}
