package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type EditDb struct {
	Db *gorm.DB
}

func NewEditDb(path string) (*EditDb, error) {
	var db, err = gorm.Open(sqlite.Open(path), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(User{})
	return &EditDb{Db: db}, nil
}
