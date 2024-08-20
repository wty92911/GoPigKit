package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/services"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		users, err := services.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.GetUser(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
