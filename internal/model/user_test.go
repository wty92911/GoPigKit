package model

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestStruct(t *testing.T) {
	validate := validator.New()
	user := User{
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
