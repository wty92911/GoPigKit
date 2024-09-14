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
		auth.POST("/family/", c.CreateFamily)
		auth.PUT("/family/join/:id", c.JoinFamily)

		auth.GET("/users", c.GetUsers)

		authFamily := auth.Group("")
		authFamily.Use(middleware.AuthFamily(false))
		{
			// 上传文件、图片，要求必须是某个家庭成员
			auth.POST("/upload_file", c.UploadFile)
			auth.DELETE("/delete_file/:url", c.DeleteFile)

			auth.GET("/categories", c.GetCategories)

			auth.GET("/all_foods", c.GetAllFoods)
			auth.GET("/foods", c.GetFoodsByCategory)
			auth.POST("/food", c.CreateFood)

			auth.GET("/menu", c.GetMenu)
			auth.POST("/menu", c.AddMenuItem)
			// 主键是(family_id, food_id)，family_id根据user可以自动推断，所以这里path参数只传food_id
			auth.PUT("/menu/:food_id", c.UpdateMenuItem)
			auth.DELETE("/menu/:food_id", c.DeleteMenuItem)

			auth.PUT("/family/exit", c.ExitFamily)

			auth.GET("/orders", c.GetOrders)
			auth.POST("/order", c.CreateOrder)
		}
		authFamily.Use(middleware.AuthFamily(true))
		{
			auth.POST("/category", c.CreateCategory)
			auth.PUT("/family/:id", c.UpdateFamily)
			auth.DELETE("/category/:id", c.DeleteCategory)
			auth.DELETE("/food/:id", c.DeleteFood)
			auth.DELETE("/order/:id", c.DeleteOrder)
			auth.DELETE("/family")

		}

	}
}
