package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
)

//TODO: 家庭级别权限隔离

// UploadFile godoc
// @Summary 上传文件
// @Description 上传文件
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "上传文件"
// @Param path formData string true "文件路径"
// @Success 200 {object} gin.H{data=string}
// @Failure 400,500 {object} gin.H{error=string}
// @Router /api/v1/upload [post]
func (ctl *Controller) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FileHeaderRequired})
		return
	}

	// 文件路径前增加用户家庭ID，便于权限隔离
	path, ok := c.GetPostForm("path")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": PathRequired})
		return
	}
	user, err := service.GetUser(c.GetString("open_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	key := fmt.Sprintf("%d/%s", user.FamilyID, path)

	// 上传文件
	key, err = service.UploadFile(fileHeader, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url := fmt.Sprintf("%s/%s/%s", database.MinIOClient.EndpointURL(), database.MinIOBucket, key)
	c.JSON(http.StatusOK, gin.H{"data": url})
}

// DeleteFile godoc
// @Summary 删除文件
// @Description 根据文件路径删除文件
// @Tags file
// @Produce json
// @Param url path string true "文件路径"
// @Success 200 {object} gin.H{message=string}
// @Failure 500 {object} gin.H{error=string}
// @Router /api/v1/file/{url} [delete]
func (ctl *Controller) DeleteFile(c *gin.Context) {
	url := c.Param("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": URLRequired})
		return
	}
	if err := service.DeleteFile(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
