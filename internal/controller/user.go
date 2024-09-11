package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
)

// GetUsers godoc
// @Summary 获得所有用户
// @Description 获取所有用户的列表
// @Tags user
// @Produce json
// @Success 200 {array} []model.User
// @Failure 500 {object} error
// @Router /api/v1/users [get]
func (ctl *Controller) GetUsers(c *gin.Context) {
	openID, exist := c.Get("openID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": OpenIDRequired})
		return
	}
	var user *model.User
	var err error
	if user, err = service.GetUser(openID.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
