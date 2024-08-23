package database__test

import (
	"github.com/wty92911/GoPigKit/configs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

func TestDatabaseInit(t *testing.T) {
	// 使用 SQLite 内存数据库进行测试，避免修改实际数据库
	var err error
	//database.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	err = database.Init(&configs.DatabaseConfig{
		Host:     "81.70.53.202",
		Port:     3306,
		User:     "pigkitadmin",
		Password: "PigkitAdmin123",
		Name:     "pigkit_test",
	})
	assert.Nil(t, err, "Database initialization should not return an error")

	// 测试数据库迁移
	err = database.DB.AutoMigrate(&model.Family{}, &model.User{}, &model.Food{}, &model.Order{}, &model.OrderItem{})
	assert.Nil(t, err, "Database migration should not return an error")
}
