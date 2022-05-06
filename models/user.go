package models

import "gorm.io/gorm"

type User struct {
	Id       uint   `gorm:"primary key;autoincrement" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}