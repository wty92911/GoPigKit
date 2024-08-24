package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetAllFamilies 获得所有家庭
func (ctl *Controller) GetAllFamilies(c *gin.Context) {
	families, err := service.GetAllFamilies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, families)
}

// GetFamilyWithPreloads 获得家庭详细情况
func (ctl *Controller) GetFamilyWithPreloads(c *gin.Context) {
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)

	// Users, Foods, Orders, Orders.Items, MenuItems
	preloads, ok := c.GetQueryArray("preloads")
	if !ok {
		preloads = []string{}
	}

	family, err := service.GetFamilyWithPreloads(uint(idInt), preloads)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, family)
}

// CreateFamily 创建家庭
func (ctl *Controller) CreateFamily(c *gin.Context) {
	openID, exist := c.Get("openID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "openID not exist"})
		return
	}
	family, err := service.CreateFamily(openID.(string), c.Query("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, family)
}

// JoinFamily 加入家庭
func (ctl *Controller) JoinFamily(c *gin.Context) {
	id := c.Query("id")
	familyID, _ := strconv.Atoi(id)
	openID, exist := c.Get("openID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "openID not exist"})
		return
	}
	family, err := service.JoinFamily(uint(familyID), openID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, family)
}
