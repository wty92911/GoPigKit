package controllers

import (
	"github.com/wty92911/GoPigKit/configs"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log("test config")
	config := configs.NewConfig()
	err := config.Update("../configs/config.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Log(config.Database.MinIO.AccessKey)
}
