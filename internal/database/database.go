package database

import (
	"fmt"
	"github.com/wty92911/GoPigKit/configs"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(config *configs.DatabaseConfig) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(
		&model.User{},
		&model.Family{},
		&model.Food{},
		&model.Order{},
		&model.OrderItem{},
	)
	return err
}
