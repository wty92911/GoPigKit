// Package controller 设置路由、解析请求字段，请求具体service，返回
package controller

import (
	"github.com/wty92911/GoPigKit/configs"
)

type Controller struct {
	Config *configs.Config
}

func NewController(config *configs.Config) *Controller {
	return &Controller{
		Config: config,
	}
}
