package controller

import (
	"github.com/wty92911/GoPigKit/internal/middleware"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"github.com/wty92911/GoPigKit/pkg/wxhelper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}
type WeChatLoginRequest struct {
	Code     string   `json:"code" binding:"required"`
	UserInfo UserInfo `json:"userInfo"`
}

// WeChatLogin 后端绑定微信登陆, 返回token
func (ctl *Controller) WeChatLogin(c *gin.Context) {
	var req WeChatLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rsp, err := wxhelper.Code2Session(ctl.Config.App, req.Code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := &model.User{
		OpenID:    rsp.OpenID,
		Name:      req.UserInfo.Nickname,
		AvatarURL: req.UserInfo.AvatarURL,
	}
	if ok := service.ExistUser(rsp.OpenID); !ok {
		err = service.CreateUser(user)
	} else {
		err = service.UpdateUser(user)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the Claims
	claims := middleware.Claims{
		OpenID: rsp.OpenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    ctl.Config.App.Name,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(ctl.Config.App.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
