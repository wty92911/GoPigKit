package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetFoodsByCategory godoc
// @Summary 根据分类获取食物
// @Description 根据分类ID获取食物列表
// @Tags food
// @Produce json
// @Param category_id query int true "分类ID"
// @Success 200 {array} {"data": []model.Food}
// @Failure 500 {object} {"error": error}
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
	c.JSON(http.StatusOK, foods)
}

// CreateFood godoc
// @Summary 创建食品
// @Description 创建食品，返回创建好的食品
// @Tags food
// @Accept json
// @Produce json
// @Param food body model.CreateFoodRequest true "创建食品请求参数"
// @Success 200 {object} model.Food
// @Failure 400 {object} map[string]interface{}{"error": "Invalid request"}
// @Failure 500 {object} map[string]interface{}{"error": "Internal server error"}
// @Router /api/v1/food [post]
func (ctl *Controller) CreateFood(c *gin.Context) {
	var req model.CreateFoodRequest

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

	c.JSON(http.StatusOK, food)
}
