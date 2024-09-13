package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetMenu godoc
// @Summary 获取菜单
// @Description 获取本家庭的菜单列表
// @Tags menu
// @Produce json
// @Success 200 {object} gin.H{data=[]model.MenuItem}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/menu [get]
func (ctl *Controller) GetMenu(c *gin.Context) {
	familyID, exist := c.Get("family_id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": FamilyIDRequired})
		return
	}
	var menuItems []*model.MenuItem
	var err error
	if menuItems, err = service.GetMenuItems(familyID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": menuItems})
}

// AddMenuItem godoc
// @Summary 添加菜单项
// @Description 添加菜单项,创建人和家庭根据token自动绑定
// @Tags menu
// @Produce json
// @Param food_id body uint true "食品ID"
// @Param quantity body uint true "数量"
// @Success 200 {object} gin.H{data=model.MenuItem}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/menu [post]
func (ctl *Controller) AddMenuItem(c *gin.Context) {
	openID := c.GetString("openid")
	if openID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": OpenIDRequired})
		return
	}
	familyID := c.GetUint("family_id")
	if familyID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": FamilyIDRequired})
		return
	}
	var menuItem model.MenuItem
	if err := c.ShouldBindJSON(&menuItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menuItem.FamilyID = &familyID
	menuItem.CreatedBy = &openID
	if err := service.AddMenuItem(&menuItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": menuItem})
}

// UpdateMenuItem godoc
// @Summary 更新菜单项
// @Description 根据path中的food_id更新菜单项，创建人和家庭根据token自动绑定，通常只会更新数量
// @Tags menu
// @Produce json
// @Param food_id path uint true "食品ID"
// @Param quantity body uint true "数量"
// @Success 200 {object} gin.H{data=model.MenuItem}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/menu/{food_id} [post]
func (ctl *Controller) UpdateMenuItem(c *gin.Context) {
	openID := c.GetString("openid")
	if openID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": OpenIDRequired})
		return
	}
	familyID := c.GetUint("family_id")
	if familyID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": FamilyIDRequired})
		return
	}
	quantity, _ := c.GetPostForm("quantity")
	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidQuantity})
		return
	}
	foodID, err := strconv.Atoi(c.Param("food_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidFoodID})
		return
	}
	uintFoodID := uint(foodID)
	menuItem := &model.MenuItem{
		FamilyID:  &familyID,
		FoodID:    &uintFoodID,
		Quantity:  uint(quantityInt),
		CreatedBy: &openID,
	}
	if err := service.UpdateMenuItem(menuItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": menuItem})
}

// DeleteMenuItem godoc
// @Summary 删除菜单项
// @Description 根据path中的food_id删除菜单项，创建人和家庭根据token自动绑定
// @Tags menu
// @Produce json
// @Param food_id path uint true "食品ID"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/menu/{food_id} [delete]
func (ctl *Controller) DeleteMenuItem(c *gin.Context) {
	openID := c.GetString("openid")
	if openID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": OpenIDRequired})
		return
	}
	familyID := c.GetUint("family_id")
	if familyID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": FamilyIDRequired})
		return
	}
	foodID, err := strconv.Atoi(c.Param("food_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidFoodID})
		return
	}
	uintFoodID := uint(foodID)
	if err := service.DeleteMenuItem(familyID, uintFoodID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
