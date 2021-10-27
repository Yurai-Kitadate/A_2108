package router

import (
	"github.com/gin-gonic/gin"
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
	r.GET("/user", con.UserGet)

	r.GET("/user/:id", controller.MockGetUserByID)
	r.GET("/plan", controller.MockGetAllPlans)
	r.GET("/plan/:id", controller.MockGetPlanByID)
	//r.GET("/plan", con.PlanGet)
	//r.GET("/plan/:id", con.PlanGetPathParam)
	r.POST("/plan", con.PlanPost)
	r.DELETE("/plan/:id", con.PlanDelete)
	r.POST("/image", con.ImagePost)
	r.POST("/register", con.RegisterPost)
	r.POST("/login", con.LoginPOST)
	return r
}
