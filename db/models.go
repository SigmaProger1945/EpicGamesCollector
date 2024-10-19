package db

type Users struct {
	email    string `gorm:"unique;not null"`
	password string `gorm:"not null"`
}
