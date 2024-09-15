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
	OpenID    string `json:"open_id"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}
type WeChatLoginRequest struct {
	Code     string   `json:"code" binding:"required"`
	UserInfo UserInfo `json:"user_info"`
}

// WeChatLogin godoc
// @Summary 后端绑定微信登陆, 返回token
// @Description 用户使用微信登录，后端绑定微信账户并返回JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param req body WeChatLoginRequest true "微信登录请求"
// @Success 200 {object} gin.H{token=string}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/login [post]
func (ctl *Controller) WeChatLogin(c *gin.Context) {
	var req WeChatLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var rsp *wxhelper.Code2SessionResponse
	var err error
	if gin.Mode() != gin.TestMode {
		rsp, err = wxhelper.Code2Session(ctl.Config.App, req.Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		// test mode
		rsp = &wxhelper.Code2SessionResponse{
			OpenID: req.UserInfo.OpenID,
		}
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
	tokenString, err := token.SignedString([]byte(ctl.Config.App.JwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token, err: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
