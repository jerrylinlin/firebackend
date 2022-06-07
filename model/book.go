package model

type Book struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Name   string `gorm:"type:varchar(60);not null;" json:"name"`
	Count  string `gorm:"type:varchar(60);not null;" json:"count"`
	Author string `gorm:"type:varchar(60);not null;" json:"author"`
	Type   string `gorm:"type:varchar(60);not null;" json:"type"`
}
