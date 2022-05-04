package models

type Book struct {
	Id            uint   `gorm:"primary key;autoincrement" json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	DatePublished string `json:"date_published"`
	Genre         string `json:"genre"`
}
