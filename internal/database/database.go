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
var MinioClient *minio.Client

func Init(config *configs.DatabaseConfig) error {
	var err error
	// Init mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("Database DSN:", dsn) // Debug print
	if err != nil {
		return err
	}

	// Init Minio
	MinioClient, err = minio.New("play.min.io", &minio.Options{
		Creds:  credentials.NewStaticV4("your-access-key", "your-secret-key", ""),
		Secure: true,
	})
	if err != nil {
		return err
	}
	//err = DB.AutoMigrate(
	//	&model.Family{},
	//	&model.User{},
	//	&model.Food{},
	//
	//	&model.Order{},
	//	&model.OrderItem{},
	//)
	return err
}
