package models

import (
	"github.com/kiyoshi-87/library-management-system/PKG/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func init() {
	DB = config.Connect()
	DB.AutoMigrate(&Book{}) //Creates a table if it doesn't exist
}

func (b *Book) CreateBook() *Book {
	DB.Create(&b)
	return b
}
