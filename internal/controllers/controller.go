// Package controllers 设置路由、解析请求字段，请求具体service，返回
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/services"
)

type Controller struct {
	service services.Service
}

func (c *Controller) SetupRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user", GetUsers)
	router.POST("/user")
}
