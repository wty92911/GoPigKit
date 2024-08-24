package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/controller"
	"github.com/wty92911/GoPigKit/internal/middleware"
)

func Init(r *gin.Engine, c *controller.Controller) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login")
	auth := r.Group("/api/v1")
	auth.Use(middleware.AuthToken(c.Config.App.JwtSecret))
	{
		auth.GET("/family", c.GetAllFamilies)
		auth.POST("/family/create", c.CreateFamily)
		auth.POST("/family/join", c.JoinFamily)

		auth.GET("/user")
		auth.POST("/user")

		authFamily := auth.Group("")
		authFamily.Use(middleware.AuthFamily(false))
		{
			auth.GET("/food")
			auth.POST("/food")

			auth.GET("/menu")
			auth.POST("/menu")

			auth.GET("/order")
			auth.POST("/order")
		}
		authFamily.Use(middleware.AuthFamily(true))
		{
			auth.DELETE("/family")
			auth.DELETE("/food")
			auth.DELETE("/order")
		}

	}
}
