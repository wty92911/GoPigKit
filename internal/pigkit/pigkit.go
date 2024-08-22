// Package pigkit is a simple web backend server
package pigkit

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/configs"
	"github.com/wty92911/GoPigKit/internal/controller"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/router"
	"log"
	"net/http"
	"time"
)

const (
	configPath = "../configs/config.yaml"
)

func Run() {
	// 1. 初始化Config
	config := configs.NewConfig()
	err := config.Update(configPath)
	if err != nil {
		panic(err)
	}

	// 2. 初始化数据库
	err = database.Init(config.Database)
	if err != nil {
		panic(err)
	}

	// 2. 注册路由
	route := gin.Default()
	ctrl := controller.NewController(config)
	router.Init(route, ctrl)

	server := &http.Server{
		Addr:           config.Server.Addr(),
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
