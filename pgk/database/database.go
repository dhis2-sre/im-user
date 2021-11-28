package database

import (
	"fmt"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ProvideDatabase(config config.Config) (*gorm.DB, error) {
	host := config.Postgresql.Host
	port := config.Postgresql.Port
	username := config.Postgresql.Username
	password := config.Postgresql.Password
	name := config.Postgresql.DatabaseName

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, name, port)

	databaseConfig := gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}

	db, err := gorm.Open(postgres.Open(dsn), &databaseConfig)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.ClusterConfiguration{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
