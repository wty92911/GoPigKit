package router__test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wty92911/GoPigKit/configs"
	"github.com/wty92911/GoPigKit/internal/controller"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"github.com/wty92911/GoPigKit/internal/router"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

const configPath = "../../configs/config.test.yaml"

func clearDatabase(db *gorm.DB) error {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS order_items")
	db.Exec("DROP TABLE IF EXISTS menu_items")
	db.Exec("DROP TABLE IF EXISTS categories")
	db.Exec("DROP TABLE IF EXISTS foods")
	db.Exec("DROP TABLE IF EXISTS orders")
	db.Exec("DROP TABLE IF EXISTS menus")

	db.Exec("DROP TABLE IF EXISTS families")
	err := db.AutoMigrate(&model.Family{}, &model.User{}, &model.Food{}, &model.Order{}, &model.OrderItem{}, &model.MenuItem{})

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	return err
}
func setupRouter(clear bool) *gin.Engine {
	// 1. 初始化Config
	config := configs.NewConfig()
	err := config.Update(configPath)
	if err != nil {
		panic(err)
	}

	// 2. 初始化数据库
	err = database.Init(config.Database)
	if err != nil {
		panic(err)
	}
	if clear {
		err = clearDatabase(database.DB)
		if err != nil {
			panic(err)
		}
	}
	// 3. 注册路由
	route := gin.Default()
	ctrl := controller.NewController(config)
	router.Init(route, ctrl)
	gin.SetMode(gin.TestMode)
	return route
}

func TestPingRoute(t *testing.T) {
	r := setupRouter(true)

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
