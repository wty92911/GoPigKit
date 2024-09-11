// Package controller 设置路由、解析请求字段，请求具体service，返回
/*
为了快速开发，大部分字段不在controller层做校验，而是在service层做，这种情况返回500
*/
package controller

import (
	"github.com/wty92911/GoPigKit/configs"
)

type Controller struct {
	Config *configs.Config
}

type ErrMsg string

const (
	InvalidParam       ErrMsg = "invalid param"
	FileHeaderRequired        = "file header is required"
	PathRequired              = "path is required"
	OpenIDRequired            = "open id is required"
)

func NewController(config *configs.Config) *Controller {
	return &Controller{
		Config: config,
	}
}
