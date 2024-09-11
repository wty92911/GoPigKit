package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/service"
	"mime/multipart"
	"net/http"
)

// UploadFile 上传文件
func (ctl *Controller) UploadFile(c *gin.Context) {
	var fileHeader *multipart.FileHeader
	var err error

	if fileHeader, err = c.FormFile("fileHeader"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fileHeader is required"})
		return
	}

	var path string
	var ok bool
	if path, ok = c.GetPostForm("path"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
		return
	}

	var url string
	if url, err = service.UploadFile(fileHeader, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, url)
}
