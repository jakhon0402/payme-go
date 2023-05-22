package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(config *DbConfig) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open(postgres.Open(config.DataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	rawDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	rawDb.SetMaxIdleConns(config.Pool.MaxIdle)
	rawDb.SetMaxOpenConns(config.Pool.MaxOpen)
	rawDb.SetConnMaxLifetime(config.Pool.MaxLifeTime)

	if config.Migrate.Enabled {
		if err := MigratePostgres(db); err != nil {
			return nil, err
		}
	}

	return db, nil

}

func MigratePostgres(db *gorm.DB) error {
	return nil
}
