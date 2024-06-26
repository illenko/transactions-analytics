package database

import (
	"github.com/illenko/analytics-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config config.AppConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Database.Url), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
