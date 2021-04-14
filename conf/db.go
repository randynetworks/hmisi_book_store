package main

import (
	"github.com/randynetworks/hmisi_book_store/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitialDB function
func InitialDB() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/hmisiBookStore?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Fail to connect database")
	}

	_ = DB.AutoMigrate(&models.Categories{})
	_ = DB.AutoMigrate(&models.Books{})
	_ = DB.AutoMigrate(&models.Users{})
	_ = DB.AutoMigrate(&models.Borrows{})
}
