package database__test

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/wty92911/GoPigKit/configs"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
	"os"
	"testing"
)

func clearDatabase(db *gorm.DB) {
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS order_items")
	db.Exec("DROP TABLE IF EXISTS menu_items")

	db.Exec("DROP TABLE IF EXISTS foods")
	db.Exec("DROP TABLE IF EXISTS orders")
	db.Exec("DROP TABLE IF EXISTS menus")

	db.Exec("DROP TABLE IF EXISTS families")
}
func TestDatabaseInit(t *testing.T) {
	var err error
	config := configs.NewConfig()
	err = config.Update("../../configs/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	err = database.Init(config.Database)
	assert.Nil(t, err, "Database initialization should not return an error")

	clearDatabase(database.DB)
	// 测试数据库迁移
	err = database.DB.AutoMigrate(&model.Family{}, &model.User{}, &model.Category{}, &model.Food{}, &model.Order{},
		&model.OrderItem{})
	assert.Nil(t, err, "Database migration should not return an error")

	// 测试MinIO上传图片
	imgPath := "./piggy.png"
	file, err := os.Open(imgPath)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}

	info, err := database.MinIOClient.PutObject(
		context.Background(),
		config.Database.MinIO.Bucket,
		"images/category/piggy.png",
		file,
		fileStat.Size(),
		minio.PutObjectOptions{ContentType: "image/png"},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

// 测试重复删除的图片
func TestDupDeleteObj(t *testing.T) {
	var err error
	config := configs.NewConfig()
	err = config.Update("../../configs/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	err = database.Init(config.Database)
	assert.Nil(t, err, "Database initialization should not return an error")

	err = database.MinIOClient.RemoveObject(
		context.Background(),
		config.Database.MinIO.Bucket,
		"images/category/piggy1.png",
		minio.RemoveObjectOptions{ForceDelete: true},
	)
	if err != nil {
		t.Fatal(err)
	}
}
