package router

import (
	"fmt"

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
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	return r
}
