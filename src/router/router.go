package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jphacks/A_2108/src/auth"
	"github.com/jphacks/A_2108/src/controller"
	"github.com/jphacks/A_2108/src/database"
)

func Route() *gin.Engine {
	r := gin.Default()
	/* con := controller.NewControllerWithYesmanRepository() */
	db, err := database.NewDatabaseHandlerWithDBName("DAWN")
	if err != nil {
		panic(err)
	}

	con := controller.NewController(db)
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
		authGroup.POST("/user", con.CreateUser)
		authGroup.PUT("/user", con.UpdateUser)
		authGroup.DELETE("/user/:id", con.DeleteUser)

		authGroup.POST("/plan", con.CreatePlan)
		authGroup.PUT("/plan", con.UpdatePlan)
		authGroup.DELETE("/plan/:id", con.DeletePlanByID)
	}

	r.GET("/user/:id", con.GetUserByID)
	r.GET("/plan", con.GetAllPlans)
	r.GET("/plan/:id", con.GetPlanByID)

	r.POST("/register", con.RegisterPost)
	r.POST("/login", con.LoginPOST)
	return r
}
