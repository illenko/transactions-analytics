package database

import (
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
	"log/slog"
)

type Migration interface {
	Execute(dir string) error
}

type migration struct {
	log *slog.Logger
	db  *gorm.DB
}

func NewMigration(log *slog.Logger, db *gorm.DB) Migration {
	return &migration{
		log: log,
		db:  db,
	}
}

func (m migration) Execute(dir string) (err error) {

	err = goose.SetDialect("postgres")
	if err != nil {
		m.log.Error("When setting database dialect")
		return
	}

	dbConnection, err := m.db.DB()

	if err != nil {
		m.log.Error("When retrieving database connection")
		return
	}

	err = goose.Up(dbConnection, dir)

	if err != nil {
		m.log.Error("When executing migration")
		return
	}

	return nil
}
