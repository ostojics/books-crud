package models

import (
	"gorm.io/gorm"
)

type Book struct {
	Id            uint   `gorm:"primary key;autoincrement" json:"id"`
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	DatePublished string `json:"date_published" validate:"required"`
	Genre         string `json:"genre" validate:"required"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
