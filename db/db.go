package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type EditDb struct {
	Db *gorm.DB
}

func NewEditDb(path string) (*EditDb, error) {
	var db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(Users{})
	return &EditDb{Db: db}, nil
}

/*func (slf EditDb) AddToDb(gameName string) {
	slf.Db.Create(&GameList{Game: gameName})
}
func (slf EditDb) DeleteFromDb(gameName string) error {
	return slf.Db.Where("game = ?", gameName).Delete(&GameList{}).Error
}*/
