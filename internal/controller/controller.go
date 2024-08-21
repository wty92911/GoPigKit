// Package controller 设置路由、解析请求字段，请求具体service，返回
package controller

import (
	"github.com/wty92911/GoPigKit/configs"
	"github.com/wty92911/GoPigKit/pkg/wxhelper"
)

type Controller struct {
	Config   *configs.Config
	wxHelper wxhelper.WxHelper
}

func NewController(config *configs.Config) *Controller {
	return &Controller{
		Config:   config,
		wxHelper: wxhelper.WxHelper{},
	}
}
