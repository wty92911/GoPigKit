// Package configs 解析配置文件，不做功能性处理
package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server   *ServerConfig   `yaml:"server"`
	Database *DatabaseConfig `yaml:"database"`
	App      *AppConfig      `yaml:"app"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type DatabaseConfig struct {
	Sql   *SqlConfig   `yaml:"sql"`
	MinIO *MinIOConfig `yaml:"minio"`
}
type SqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
type MinIOConfig struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Bucket    string `yaml:"bucket"`
}
type AppConfig struct {
	ID        string `yaml:"id"`
	Name      string `yaml:"name"`
	Secret    string `yaml:"secret"`
	JwtSecret string `yaml:"jwt_secret"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Update(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
