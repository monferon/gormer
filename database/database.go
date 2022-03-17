package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gormer/entity"
	"log"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.Book{})

	DB = db
}
