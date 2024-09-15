// Package router 路由层，负责路由的注册、保护逻辑。
package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wty92911/GoPigKit/docs"
	"github.com/wty92911/GoPigKit/internal/controller"
	"github.com/wty92911/GoPigKit/internal/middleware"
)

// Init 初始化路由
func Init(r *gin.Engine, c *controller.Controller) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// create user or update user when login
	r.POST("/login", c.WeChatLogin)
	auth := r.Group("/api/v1")
	auth.Use(middleware.AuthToken(c.Config.App.JwtSecret))
	{
		auth.GET("/family", c.GetAllFamilies)
		auth.POST("/family/create", c.CreateFamily)
		auth.PUT("/family/join/:id", c.JoinFamily)

		auth.GET("/users", c.GetUsers)

		authFamily := auth.Group("")
		authFamily.Use(middleware.AuthFamily(false))
		{
			// 上传文件、图片，要求必须是某个家庭成员
			authFamily.POST("/file", c.UploadFile)
			// 因为是传url，所以用POST方式和upload区分开
			authFamily.POST("/file/delete", c.DeleteFile)

			authFamily.GET("/categories", c.GetCategories)

			authFamily.GET("/all_foods", c.GetAllFoods)
			authFamily.GET("/foods", c.GetFoodsByCategory)
			authFamily.POST("/food", c.CreateFood)

			authFamily.GET("/menu", c.GetMenu)
			authFamily.POST("/menu", c.AddMenuItem)
			// 主键是(family_id, food_id)，family_id根据user可以自动推断，所以这里path参数只传food_id
			authFamily.PUT("/menu/:food_id", c.UpdateMenuItem)
			authFamily.DELETE("/menu/:food_id", c.DeleteMenuItem)

			authFamily.PUT("/family/exit", c.ExitFamily)

			authFamily.GET("/orders", c.GetOrders)
			authFamily.POST("/order", c.CreateOrder)
		}
		authFamilyOwner := auth.Group("")
		authFamilyOwner.Use(middleware.AuthFamily(true))
		{
			authFamilyOwner.POST("/category", c.CreateCategory)
			authFamilyOwner.PUT("/family/:id", c.UpdateFamily)
			authFamilyOwner.DELETE("/category/:id", c.DeleteCategory)
			authFamilyOwner.DELETE("/food/:id", c.DeleteFood)
			authFamilyOwner.DELETE("/order/:id", c.DeleteOrder)
			authFamilyOwner.DELETE("/family")

		}

	}
}
