package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/controller"
	"github.com/wty92911/GoPigKit/internal/middlewares"
)

func Init(r *gin.Engine, c *controller.Controller) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login")
	auth := r.Group("/")
	auth.Use(middlewares.AuthToken(c.Config.App.JwtSecret))
	{
		auth.GET("/family")
		auth.POST("/family")
		auth.DELETE("/family")

		auth.GET("/user")
		auth.POST("/user")

		auth.GET("/food")
		auth.POST("/food")
		auth.DELETE("/food")
	}
}
