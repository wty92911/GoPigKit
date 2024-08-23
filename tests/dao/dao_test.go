package dao

import (
	"github.com/stretchr/testify/assert"
	"github.com/wty92911/GoPigKit/configs"
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
	"testing"
)

func Init() error {
	return database.Init(&configs.DatabaseConfig{
		Host:     "81.70.53.202",
		Port:     3306,
		User:     "root",
		Password: "PigkitAdmin123",
		Name:     "pigkit_test",
	})
}
func clearDatabase(db *gorm.DB) {
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS order_items")
	db.Exec("DROP TABLE IF EXISTS menu_items")

	db.Exec("DROP TABLE IF EXISTS foods")
	db.Exec("DROP TABLE IF EXISTS orders")
	db.Exec("DROP TABLE IF EXISTS menus")

	db.Exec("DROP TABLE IF EXISTS families")
}
func SameUsers(a, b []model.User) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].OpenID != b[i].OpenID {
			return false
		}
	}
	return true
}
func SameFoods(a, b []model.Food) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Title != b[i].Title {
			return false
		}
	}
	return true
}

func TestUserFoodFamily(t *testing.T) {

	err := Init()
	assert.Nil(t, err, "Database initialization should not return an error")

	clearDatabase(database.DB)

	err = database.DB.AutoMigrate(&model.Family{}, &model.User{}, &model.Food{})
	assert.Nil(t, err, "Database migration should not return an error")

	users := []model.User{
		{OpenID: "test1", Name: "test1", FamilyID: 1},
		{OpenID: "test2", Name: "test2", FamilyID: 1},
		{OpenID: "test3", Name: "test3", FamilyID: 2},
		{OpenID: "test4", Name: "test4", FamilyID: 2},
	}
	foods := []model.Food{
		{Title: "test1", Price: 100, Desc: "This is a food", Images: "[https://www.baidu.com]", FamilyID: 1},
		{Title: "test2", Price: 120, Desc: "This is a food", Images: "[https://www.baidu.com]", FamilyID: 1},
		{Title: "test3", Price: 130, Desc: "This is a food", Images: "[https://www.baidu.com]", FamilyID: 2},
		{Title: "test4", Price: 140, Desc: "This is a food", Images: "[https://www.baidu.com]", FamilyID: 2},
	}
	families := []model.Family{
		{Name: "test1", OwnerID: 1},
		{Name: "test2", OwnerID: 2},
	}
	familiesWithPreloads := []model.Family{
		{Name: "test1", OwnerID: 1, Users: []model.User{users[0], users[1]}, Foods: []model.Food{foods[0], foods[1]}},
		{Name: "test2", OwnerID: 2, Users: []model.User{users[2], users[3]}, Foods: []model.Food{foods[2], foods[3]}},
	}
	for _, family := range families {
		err = dao.CreateFamily(&family)
		assert.Nil(t, err, "Create family should not return an error")
	}
	for _, user := range users {
		err = dao.CreateUser(&user)
		assert.Nil(t, err, "Create user should not return an error")
	}
	for _, user := range users {
		getUser, err := dao.GetUser(user.OpenID)
		assert.Nil(t, err, "Get user should not return an error")
		assert.Equal(t, user.Name, getUser.Name, "Get user should return the same user")
	}

	for _, food := range foods {
		err = dao.CreateFood(&food)
		assert.Nil(t, err, "Create food should not return an error")
	}
	for i, food := range foods {
		getFood, err := dao.GetFood(uint(i + 1))
		assert.Nil(t, err, "Get food should not return an error")
		assert.Equal(t, food.Title, getFood.Title, "Get food should return the same food")
	}

	for i, family := range families {
		getFamily, err := dao.GetFamily(uint(i + 1))
		assert.Nil(t, err, "Get family should not return an error")
		assert.Equal(t, family.Name, getFamily.Name, "Get family should return the same family")
	}
	for i, family := range familiesWithPreloads {
		getFamily, err := dao.GetFamilyWithPreloads(uint(i+1), []string{"Users", "Foods"})
		assert.Nil(t, err, "Get family with preloads should not return an error")
		assert.Equal(t, SameUsers(family.Users, getFamily.Users), true, "Get family with preloads should return the same family")
		assert.Equal(t, SameFoods(family.Foods, getFamily.Foods), true, "Get family with preloads should return the same family")
	}

	user, _ := dao.GetUser("test1")
	user.FamilyID = 2
	err = dao.UpdateUser(user)
	assert.Nil(t, err, "Update user should not return an error")
	getFamilyWithPreloads, err := dao.GetFamilyWithPreloads(2, []string{"Users", "Foods"})
	assert.Nil(t, err, "Get family with preloads should not return an error")
	assert.Equal(t, SameUsers([]model.User{*user, users[2], users[3]}, getFamilyWithPreloads.Users), true, "Get family with preloads should return the same family")

	food, _ := dao.GetFood(1)
	food.Title = "test11"
	err = dao.UpdateFood(food)
	assert.Nil(t, err, "Update food should not return an error")

	family, _ := dao.GetFamily(1)
	family.Name = "test11"
	err = dao.UpdateFamily(family)
	assert.Nil(t, err, "Update family should not return an error")

	for _, user := range users {
		err := dao.DeleteUser(user.OpenID)
		assert.Nil(t, err, "Get user should not return an error")
	}
	for i := range foods {
		err := dao.DeleteFood(uint(i + 1))
		assert.Nil(t, err, "Get food should not return an error")
	}
	for i := range familiesWithPreloads {
		err := dao.DeleteFamily(uint(i + 1))
		assert.Nil(t, err, "Get family should not return an error")
	}
}

//func TestMenuOrder(t *testing.T) {
//	err := Init()
//	assert.Nil(t, err, "Database initialization should not return an error")
//
//	clearDatabase(database.DB)
//
//	err = database.DB.AutoMigrate(&model.Family{}, &model.Order{}, &model.MenuItem{}, &model.OrderItem{})
//	assert.Nil(t, err, "Database migration should not return an error")
//
//	families := []model.Family{
//		{Name: "test1", OwnerID: 1},
//		{Name: "test2", OwnerID: 2},
//	}
//	Orders := []model.Order{
//		{FamilyID: 1, Items: []model.OrderItem{{FoodID: 1, Quantity: 1}}},},
//    }
//	}
//}
