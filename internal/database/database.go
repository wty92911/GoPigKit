// Package database is a database package of gorm mysql and minio client.
package database

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/wty92911/GoPigKit/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var MinIOClient *minio.Client
var MinIOBucket string

func Init(config *configs.DatabaseConfig) error {
	var err error
	// Init mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Sql.User,
		config.Sql.Password,
		config.Sql.Host,
		config.Sql.Port,
		config.Sql.Name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("Database DSN:", dsn) // Debug print
	if err != nil {
		return err
	}

	// Init Minio
	MinIOClient, err = minio.New(config.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinIO.AccessKey, config.MinIO.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	MinIOBucket = config.MinIO.Bucket
	return err
}
