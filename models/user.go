package models

import "gorm.io/gorm"

type User struct {
	Id       uint   `gorm:"primary key;autoincrement" json:"id"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}