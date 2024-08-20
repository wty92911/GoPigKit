// Package pigkit is a simple web backend server
package pigkit

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/controllers"
)

func Run() {
	r := gin.Default()
	controllers.SetupRoutes(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
