package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetAllFamilies godoc
// @Summary 获得所有家庭
// @Description 获取所有家庭的列表
// @Tags family
// @Produce json
// @Success 200 {array} model.Family
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/family [get]
func (ctl *Controller) GetAllFamilies(c *gin.Context) {
	families, err := service.GetAllFamilies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, families)
}

// GetFamilyWithPreloads godoc
// @Summary 获得家庭详细情况
// @Description 获取指定家庭的详细信息，包括可选的预加载项
// @Tags family
// @Produce json
// @Param id query int true "家庭ID"
// @Param preloads query []string false "预加载项"
// @Success 200 {object} model.Family
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/family/details [get]
func (ctl *Controller) GetFamilyWithPreloads(c *gin.Context) {
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)

	// Users, Foods, Orders, Orders.Items, MenuItems
	var preloads []string
	preloads, ok := c.GetQueryArray("preloads")
	if !ok {
		preloads = []string{}
	}
	var family *model.Family
	family, err := service.GetFamilyWithPreloads(uint(idInt), preloads)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, family)
}

// CreateFamily godoc
// @Summary 创建家庭
// @Description 创建一个新的家庭
// @Tags family
// @Produce json
// @Param name query string true "家庭名称"
// @Success 200 {object} model.Family
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/family/create [post]
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

// JoinFamily godoc
// @Summary 加入家庭
// @Description 加入一个现有的家庭
// @Tags family
// @Produce json
// @Param id query int true "家庭ID"
// @Success 200 {object} model.Family
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/family/join [post]
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
