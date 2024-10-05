package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type EditDb struct {
	Db *gorm.DB
}

func NewEditDb() (*EditDb, error) {
	var db, err = gorm.Open(sqlite.Open("db/db.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&GameList{})
	return &EditDb{Db: db}, nil
}

func (slf EditDb) AddGame(gameName string) {
	slf.Db.Create(&GameList{Game: gameName})
}
func (slf EditDb) DeleteProxy(gameName string) error {
	return slf.Db.Where("game = ?", gameName).Delete(&GameList{}).Error
}
