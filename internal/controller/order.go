package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
	"strconv"
)

// GetOrders godoc
// @Summary 获取订单列表
// @Description 获取用户家庭的订单列表
// @Tags order
// @Produce json
// @Success 200 {object} gin.H{data=[]model.Order}
// @Failure 500 {object} gin.H{error=string}
// @Router /api/v1/orders [get]
func (ctl *Controller) GetOrders(c *gin.Context) {
	familyID := c.GetUint("family_id")
	family, err := service.GetFamilyWithPreloads(familyID, []string{"Orders.Items"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": family.Orders})
}

// CreateOrder godoc
// @Summary 创建订单
// @Description 创建订单
// @Tags order
// @Accept json
// @Produce json
// @Param items body []model.MenuItem true "订单项"
// @Success 200 {object} gin.H{message=string}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/orders [post]
func (ctl *Controller) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	familyID := c.GetUint("family_id")
	order.FamilyID = &familyID

	// item is a pointer, so we can set it's attr
	openID := c.GetString("open_id")
	for _, item := range order.Items {
		if item.CreatedBy == nil {
			item.CreatedBy = &openID
		}
	}
	err := service.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// DeleteOrder godoc
// @Summary 删除订单
// @Description 删除订单
// @Tags 订单
// @Accept  json
// @Produce  json
// @Param id path uint true "订单ID"
// @Success 200 {object} gin.H{message=string}
// @Failure 500 {object} gin.H{error=string}
// @Router /api/v1/order/{id} [delete]
func (ctl *Controller) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	err := service.DeleteOrder(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
