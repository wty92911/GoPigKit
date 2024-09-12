package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetAllFoods godoc
// @Summary 获得所有食物
// @Description 获取自己家庭的所有食物的列表
// @Tags food
// @Produce json
// @Success 200 {object} gin.H{data=[]model.Food}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/food [get]
func (ctl *Controller) GetAllFoods(c *gin.Context) {
	familyID, exist := c.Get("family_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "family_id not exist"})
		return
	}
	var foods []*model.Food
	var err error
	if foods, err = service.GetAllFoods(familyID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foods})
}

// GetFoodsByCategory godoc
// @Summary 根据分类获取食物
// @Description 根据分类ID获取食物列表
// @Tags food
// @Produce json
// @Param category_id query int true "分类ID"
// @Success 200 {object} gin.H{data=[]model.Food}
// @Failure 500 {object} gin.H{error=string}
// @Router /api/v1/foods [get]
func (ctl *Controller) GetFoodsByCategory(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Query("category_id"))
	var foods []*model.Food
	var err error
	foods, err = service.GetFoodsByCategoryID(uint(categoryID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foods})
}

// CreateFood godoc
// @Summary 创建食品
// @Description 创建食品，返回创建好的食品
// @Tags food
// @Accept json
// @Produce json
// @Param food body model.Food true "创建食品请求参数"
// @Success 200 {object} gin.H{data=model.Food}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/food [post]
func (ctl *Controller) CreateFood(c *gin.Context) {
	var req *model.Food
	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层创建食品
	food, err := service.CreateFood(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": food})
}
