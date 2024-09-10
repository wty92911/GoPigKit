package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/GoPigKit/internal/service"
	"mime/multipart"
	"net/http"
)

func (ctl *Controller) UploadFile(c *gin.Context) {
	var fileHeader *multipart.FileHeader
	var err error

	if fileHeader, err = c.FormFile("fileHeader"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fileHeader is required"})
		return
	}

	var file multipart.File
	file, err = fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var path string
	var ok bool
	if path, ok = c.GetPostForm("path"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
		return
	}

	if path, err = service.UploadFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, path)
}
