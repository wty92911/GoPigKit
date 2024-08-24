package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
)

func (ctl *Controller) GetUsers(c *gin.Context) {
	openID, exist := c.Get("openID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "openID not exist"})
		return
	}
	user, err := service.GetUser(openID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
