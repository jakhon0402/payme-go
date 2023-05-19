package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"payme-go/internal/config"
)

func OpenPostgres(config *config.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open(postgres.Open(config.DB.DataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	rawDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	rawDb.SetMaxIdleConns(config.DB.Pool.MaxIdle)
	rawDb.SetMaxOpenConns(config.DB.Pool.MaxOpen)
	rawDb.SetConnMaxLifetime(config.DB.Pool.MaxLifeTime)

	if config.DB.Migrate.Enabled {
		if err := MigratePostgres(db); err != nil {
			return nil, err
		}
	}

	return db, nil

}

func MigratePostgres(db *gorm.DB) error {
	return nil
}
