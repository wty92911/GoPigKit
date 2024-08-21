package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/wty92911/GoPigKit/internal/model"
	"testing"
)

func TestStruct(t *testing.T) {
	validate := validator.New()
	user := model.User{
		Name: "jacklove",
		Role: "diner",
	}
	err := validate.Struct(user)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}
