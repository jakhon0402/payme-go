package config

import (
	"github.com/spf13/viper"
	"payme-go/pkg/database"
)

type ServerConfig struct {
	Port string `json:"port" yaml:"port"`
}

type Config struct {
	Server ServerConfig      `json:"server" yaml:"server"`
	DB     database.DbConfig `json:"db" yaml:"db"`
}

func LoadConfig() (*Config, error) {
	vp := viper.New()
	var config Config
	vp.AddConfigPath("././config/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := vp.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
