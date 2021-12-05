package storage

import (
	"fmt"
	"github.com/dhis2-sre/im-users/pkg/config"
	"github.com/dhis2-sre/im-users/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ProvideDatabase(c config.Config) (*gorm.DB, error) {
	host := c.Postgresql.Host
	port := c.Postgresql.Port
	username := c.Postgresql.Username
	password := c.Postgresql.Password
	name := c.Postgresql.DatabaseName

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
