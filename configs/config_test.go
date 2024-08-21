package configs

import "testing"

func TestConfig(t *testing.T) {
	t.Log("test config")
	config := NewConfig()
	err := config.Update("./config.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Log(config)
}
