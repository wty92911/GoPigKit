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
// @Success 200 {object} gin.H{data=[]model.Family}
// @Failure 500 {object} gin.H{error=string}
// @Router /api/v1/family [get]
func (ctl *Controller) GetAllFamilies(c *gin.Context) {
	families, err := service.GetAllFamilies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": families})
}

// GetFamilyWithPreloads godoc
// @Summary 获得家庭详细情况
// @Description 获取指定家庭的详细信息，包括可选的预加载项
// @Tags family
// @Produce json
// @Param id query int true "家庭ID"
// @Param preloads query []string false "预加载项，如:Users, Foods, Orders, Orders.Items, MenuItems"
// @Success 200 {object} gin.H{data=model.Family}
// @Failure 500 {object} gin.H{error=string}
// @Router /api/v1/family/details [get]
func (ctl *Controller) GetFamilyWithPreloads(c *gin.Context) {
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)

	preloads, ok := c.GetQueryArray("preloads")
	if !ok {
		preloads = []string{}
	}
	family, err := service.GetFamilyWithPreloads(uint(idInt), preloads)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": family})
}

// CreateFamily godoc
// @Summary 创建家庭
// @Description 创建一个新的家庭
// @Tags family
// @Produce json
// @Param name query string true "家庭名称"
// @Success 200 {object} gin.H{data=model.Family}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/family/create [post]
func (ctl *Controller) CreateFamily(c *gin.Context) {
	openID := c.GetString("openid")
	if openID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": OpenIDRequired})
		return
	}
	var family *model.Family
	var err error
	family, err = service.CreateFamily(openID, c.Query("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": family})
}

// UpdateFamily godoc
// @Summary 更新家庭
// @Description 更新家庭
// @Tags family
// @Accept json
// @Produce json
// @Param name body string true "家庭名称"
// @Param owner_open_id body string false "家庭owner的openID"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/family [put]
func (ctl *Controller) UpdateFamily(c *gin.Context) {
	var req *model.Family
	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	openID := c.GetString("openid")
	user, err := service.GetUser(openID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	family, err := service.GetFamilyWithPreloads(*user.FamilyID, []string{})

	// update family
	family.Name = req.Name
	if req.OwnerOpenID != nil {
		family.OwnerOpenID = req.OwnerOpenID
	}
	err = service.UpdateFamily(family)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// JoinFamily godoc
// @Summary 加入家庭
// @Description 加入一个现有的家庭
// @Tags family
// @Produce json
// @Param id path int true "家庭ID"
// @Success 200 {object} gin.H{data=model.Family}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/family/join/{id} [put]
func (ctl *Controller) JoinFamily(c *gin.Context) {
	id := c.Param("id")
	familyID, _ := strconv.Atoi(id)
	openID, exist := c.Get("openid")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": OpenIDRequired})
		return
	}
	family, err := service.JoinFamily(uint(familyID), openID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": family})
}

// ExitFamily godoc
// @Summary 退出家庭
// @Description 退出当前家庭
// @Tags family
// @Produce json
// @Success 200 {object} gin.H{message=string}
// @Failure 500 {object} gin.H{error=string}
// @Router /api/v1/family/exit [put]
func (ctl *Controller) ExitFamily(c *gin.Context) {
	openID := c.GetString("openid")
	if err := service.ExitFamily(openID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
