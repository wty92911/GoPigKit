// Package controller (handler) 接受请求，解析请求字段，请求具体service，返回
/*
为了快速开发，大部分字段不在controller层做校验，而是在service层被动做，这种情况返回500
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
	InvalidParam          ErrMsg = "invalid param"
	FileHeaderRequired           = "file header is required"
	PathRequired                 = "path is required"
	URLRequired                  = "url is required"
	OpenIDRequired               = "open id is required"
	NameRequired                 = "name is required"
	FamilyIDRequired             = "family id is required"
	QuantityRequired             = "quantity is required"
	InvalidFoodID                = "invalid food id"
	InvalidQuantity              = "invalid quantity"
	FamilyOwnerCannotExit        = "family owner cannot exit"
)

func NewController(config *configs.Config) *Controller {
	return &Controller{
		Config: config,
	}
}
