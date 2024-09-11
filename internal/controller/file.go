package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/service"
	"net/http"
)

// UploadFile godoc
// @Summary 上传文件
// @Description 上传文件
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "上传文件"
// @Param path formData string true "文件路径"
// @Success 200 {string} string "url"
// @Failure 400 {object} ErrMsg
// @Failure 500 {object} error
// @Router /api/v1/upload [post]
func (ctl *Controller) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": FileHeaderRequired})
		return
	}

	path, ok := c.GetPostForm("path")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": PathRequired})
		return
	}

	url, err := service.UploadFile(fileHeader, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, url)
}

// DeleteFile 删除文件
func (ctl *Controller) DeleteFile(c *gin.Context) {
	var path string
	var ok bool
	if path, ok = c.GetPostForm("path"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": PathRequired})
		return
	}

	if err := service.DeleteFile(path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
