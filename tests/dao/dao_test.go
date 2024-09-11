package dao

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/wty92911/GoPigKit/configs"
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
	"testing"
)

func Init() error {
	config := configs.NewConfig()
	err := config.Update("../../configs/config.yaml")
	if err != nil {
		return err
	}
	return database.Init(config.Database)

}
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
func SameOrderItems(a, b []model.OrderItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].OrderID != b[i].OrderID || a[i].FoodID != b[i].FoodID || a[i].Quantity != b[i].Quantity {
			return false
		}
	}
	return true
}

func SameMenuItems(a, b []model.MenuItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].FamilyID != b[i].FamilyID || a[i].FoodID != b[i].FoodID || a[i].Quantity != b[i].Quantity {
			return false
		}
	}
	return true
}

// CompareFamilies 比较两个 Family 结构体是否相等
func CompareFamilies(f1, f2 *model.Family) bool {
	opts := []cmp.Option{cmpopts.IgnoreFields(model.Family{},
		"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Owner", "Users", "Orders", "MenuItems"),
		cmpopts.IgnoreFields(model.Category{},
			"ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
		cmpopts.IgnoreFields(model.Food{},
			"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "CreatedUser")}
	fmt.Println(cmp.Diff(*f1, *f2, opts...))
	return cmp.Equal(*f1, *f2, opts...)

}
func TestAll(t *testing.T) {

	err := Init()
	assert.Nil(t, err, "Database initialization should not return an error")

	err = clearDatabase(database.DB)
	assert.Nil(t, err, "Database clear should not return an error")

	users := []*model.User{
		{OpenID: "test1", Name: "test1", FamilyID: nil},
		{OpenID: "test2", Name: "test2", FamilyID: nil},
		{OpenID: "test3", Name: "test3", FamilyID: nil},
		{OpenID: "test4", Name: "test4", FamilyID: nil},
	}
	// create users
	for i, user := range users {
		err = dao.CreateUser(database.DB, user)
		assert.Nil(t, err, "Create user should not return an error")
		users[i], err = dao.GetUser(user.OpenID)
		assert.Nil(t, err, "Get user should not return an error")
	}
	families := []*model.Family{
		{Name: "test1", OwnerOpenID: &users[0].OpenID},
		{Name: "test2", OwnerOpenID: &users[2].OpenID},
	}
	// create families
	for i, family := range families {
		err = dao.CreateFamily(database.DB, family)
		assert.Nil(t, err, "Create family should not return an error")
		families[i], err = dao.GetFamily(family.ID)
		assert.Nil(t, err, "Get family should not return an error")
	}

	// user join a family
	for i := range users {
		family, err := dao.GetFamily(uint(i/2 + 1))
		assert.Nil(t, err, "Get family should not return an error")
		users[i].FamilyID = &family.ID
		err = dao.UpdateUser(database.DB, users[i])
		assert.Nil(t, err, "Update user should not return an error")
	}
	categories := []*model.Category{
		{FamilyID: &families[0].ID, TopName: "TopName", MidName: "MidName", Name: "Name1", ImageURL: "https://www.baidu.com", Foods: []*model.Food{}},
		{FamilyID: &families[0].ID, TopName: "TopName", MidName: "MidName", Name: "Name2", ImageURL: "https://www.baidu.com", Foods: []*model.Food{}},
		{FamilyID: &families[1].ID, TopName: "TopName", MidName: "MidName", Name: "Name1", ImageURL: "https://www.baidu.com", Foods: []*model.Food{}},
		{FamilyID: &families[1].ID, TopName: "TopName", MidName: "MidName", Name: "Name2", ImageURL: "https://www.baidu.com", Foods: []*model.Food{}},
	}
	// create categories
	for i, category := range categories {
		err = dao.CreateCategory(database.DB, category)
		assert.Nil(t, err, "Create category should not return an error")
		categories[i], err = dao.GetCategory(category.ID)
		assert.Nil(t, err, "Get category should not return an error")
	}

	foods := []*model.Food{
		{Title: "test1", Price: 100, Desc: "This is a food", ImageURLs: json.RawMessage(`[https://www.baidu.com]`), CategoryID: &categories[0].ID, CreatedBy: &users[0].OpenID},
		{Title: "test2", Price: 120, Desc: "This is a food", ImageURLs: json.RawMessage(`[https://www.baidu.com]`), CategoryID: &categories[0].ID, CreatedBy: &users[1].OpenID},
		{Title: "test3", Price: 130, Desc: "This is a food", ImageURLs: json.RawMessage(`[https://www.baidu.com]`), CategoryID: &categories[2].ID, CreatedBy: &users[2].OpenID},
		{Title: "test4", Price: 140, Desc: "This is a food", ImageURLs: json.RawMessage(`[https://www.baidu.com]`), CategoryID: &categories[2].ID, CreatedBy: &users[3].OpenID},
	}
	//create foods
	for i, food := range foods {
		err = dao.CreateFood(database.DB, food)
		assert.Nil(t, err, "Create food should not return an error")
		foods[i], err = dao.GetFood(food.ID)
		assert.Nil(t, err, "Get food should not return an error")
	}

	// get categories with foods
	for i := range categories {
		categories[i], err = dao.GetCategoryWithPreloads(uint(i+1), []string{"Foods"})
		assert.Nil(t, err, "Get category should not return an error")
	}
	familiesWithPreloads := []*model.Family{
		{Name: "test1", OwnerOpenID: &users[0].OpenID, Users: []*model.User{users[0], users[1]}, Categories: []*model.Category{categories[0], categories[1]}},
		{Name: "test2", OwnerOpenID: &users[2].OpenID, Users: []*model.User{users[2], users[3]}, Categories: []*model.Category{categories[2], categories[3]}},
	}

	for i, family := range families {
		getFamily, err := dao.GetFamily(uint(i + 1))
		assert.Nil(t, err, "Get family should not return an error")
		assert.Equal(t, family.Name, getFamily.Name, "Get family should return the same family")
	}
	for i, family := range familiesWithPreloads {
		getFamily, err := dao.GetFamilyWithPreloads(uint(i+1), []string{"Users", "Categories.Foods"})
		assert.Nil(t, err, "Get family with preloads should not return an error")
		assert.True(t, CompareFamilies(family, getFamily), "Get family with preloads should return the same family")
	}
	for i := range foods {
		err := dao.DeleteFood(database.DB, uint(i+1))
		assert.Nil(t, err, "Get food should not return an error")
	}
	for i := range categories {
		err := dao.DeleteCategory(database.DB, uint(i+1))
		assert.Nil(t, err, "Get category should not return an error")
	}
	// exit family
	for i := range users {
		getUser, err := dao.GetUser(users[i].OpenID)
		assert.Nil(t, err, "Get user should not return an error")
		getUser.FamilyID = nil
		err = dao.UpdateUser(database.DB, getUser)
		assert.Nil(t, err, "Update user should not return an error")
	}
	for i := range familiesWithPreloads {
		err := dao.DeleteFamily(database.DB, uint(i+1))
		assert.Nil(t, err, "Get family should not return an error")
	}
	for i := range users {
		err := dao.DeleteUser(database.DB, users[i].OpenID)
		assert.Nil(t, err, "Get user should not return an error")
	}
}

func TestCreateNullForeignKey(t *testing.T) {
	err := Init()
	assert.Nil(t, err, "Database initialization should not return an error")

	err = clearDatabase(database.DB)
	assert.Nil(t, err, "Database clear should not return an error")

	user := model.User{
		OpenID: "test1", Name: "test",
	}
	err = dao.CreateUser(database.DB, &user)
	assert.Nil(t, err, "Create user should not return an error")

	user1, _ := dao.GetUser("test1")
	assert.Equal(t, user1.FamilyID, 0)
}
