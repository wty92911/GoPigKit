package service

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/wty92911/GoPigKit/internal/database"
	"log"
	"mime/multipart"
	"strings"
)

// UploadFile 上传文件，返回文件Key
func UploadFile(fileHeader *multipart.FileHeader, key string) (string, error) {
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
		key,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{ContentType: "image/png"},
	)
	if err != nil {
		return "", err
	}
	return info.Key, nil
}

// DeleteFile 根据path删除文件
// url格式：http://127.0.0.1:9000/GoPigKit/1619160061.png，
func DeleteFile(path string) error {
	// 找到真实的path
	path := strings.TrimPrefix(url,
		fmt.Sprintf("%s/%s/", database.MinIOClient.EndpointURL(), database.MinIOBucket))
	err := database.MinIOClient.RemoveObject(
		context.Background(),
		database.MinIOBucket,
		path,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("delete file %s error: %v, path is %s", url, err, path)
	}
	return nil
}
