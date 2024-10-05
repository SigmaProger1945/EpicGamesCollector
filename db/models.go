package db

type GameList struct {
	Game string `gorm:"<-";unique;not null`
}

func (GameList) TableName() string {
	return "glist"
}
