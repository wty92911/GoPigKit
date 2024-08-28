package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetCategories godoc
// @Summary 获得所有分类
// @Description 获取所有分类的列表
// @Tags category
// @Produce json
// @Success 200 {array} model.Category
// @Failure 500 {object} map[string]interface{}
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

func (ctl *Controller) CreateCategory(c *gin.Context) {
	id, _ := c.GetPostForm("family_id")
	familyID, _ := strconv.Atoi(id)
	topName, _ := c.GetPostForm("top_name")
	midName, _ := c.GetPostForm("mid_name")
	name, _ := c.GetPostForm("name")
	// Read image file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err := service.CreateCategory(uint(familyID), topName, midName, name, file)
	// Upload image file

}
