package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetCategories godoc
// @Summary 获得所有分类
// @Description 获取所有分类的列表
// @Tags category
// @Produce json
// @Param family_id query uint true "家庭ID"
// @Success 200 {array} {"data": []model.Category}
// @Failure 500 {object} {"error": error}
// @Router /api/v1/categories [get]
func (ctl *Controller) GetCategories(c *gin.Context) {
	familyID, _ := strconv.Atoi(c.Query("family_id"))
	categories, err := service.GetCategories(uint(familyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary 创建分类
// @Description 创建分类,包括分类的三级名称、图片，返回创建好的分类+图片链接
// @Tags category
// @Accept multipart/form-data
// @Produce json
// @Param family_id body uint true "家庭ID"
// @Param top_name body string true "顶级分类名称"
// @Param mid_name body string true "中间分类名称"
// @Param name body string true "分类名称"
// @Param image_url body string true "图片链接"
// @Success 200 {object} {"data": model.Category}
// @Failure 400 {object} {"error": error}
// @Failure 500 {object} {"error": error}
// @Router /api/v1/category [post]
func (ctl *Controller) CreateCategory(c *gin.Context) {
	var req *model.Category
	// 绑定并验证请求参数
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := service.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary 删除分类
// @Description 根据删除分类
// @Tags category
// @Produce json
// @Param id path int true "分类ID"
// @Success 200 {object} {"message": "success"}
// @Failure 500 {object} {"error": error}
// @Router /api/v1/category/{id} [delete]
func (ctl *Controller) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
