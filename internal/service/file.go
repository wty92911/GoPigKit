package service

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/wty92911/GoPigKit/internal/database"
	"log"
	"mime/multipart"
)

// UploadFile 上传文件，返回文件真实路径
func UploadFile(fileHeader *multipart.FileHeader, path string) (string, error) {
	file, err := fileHeader.Open()
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	if err != nil {
		return "", err
	}

	info, err := database.MinIOClient.PutObject(
		context.Background(),
		database.MinIOBucket,
		path,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{ContentType: "image/png"},
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", database.MinIOClient.EndpointURL(), info.Bucket, info.Key), nil
}

// DeleteFile 删除文件
func DeleteFile(path string) error {
	err := database.MinIOClient.RemoveObject(
		context.Background(),
		database.MinIOBucket,
		path,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("delete file %s error: %v", path, err)
	}
	return nil
}
