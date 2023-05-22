package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	UnsupportedDriver = errors.New("unsupported database driver")
)

type DbConfig struct {
	Driver         string `json:"driver" json:"driver"`
	DataSourceName string `yaml:"data_source_name" json:"data_source_name"`
	Migrate        struct {
		Enabled bool `json:"enabled" yaml:"enabled"`
		//Dir     string `json:"dir" yaml:"dir"`
	} `json:"migrate" yaml:"migrate"`
	Pool struct {
		MaxOpen     int           `json:"max-open" yaml:"max-open"`
		MaxIdle     int           `json:"max-idle" yaml:"max-idle"`
		MaxLifeTime time.Duration `json:"max-lifetime" yaml:"max-lifetime"`
	} `json:"pool" yaml:"pool"`
}

func OpenDB(config *DbConfig) (*gorm.DB, error) {
	switch config.Driver {
	case "postgres":
		return OpenPostgres(config)
	default:
		return nil, UnsupportedDriver
	}
}
