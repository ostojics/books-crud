package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id            uint   `gorm:"primary key;autoincrement" json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	DatePublished time.Time `json:"date_published"`
	Genre         string `json:"genre"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
