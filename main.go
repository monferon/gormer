package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     int
}

func (u User) TableName() string {
	return "User"
}

func main() {
	var User User
	dsn := "host=localhost user=postgres password=qwerty dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//firstUser := User{}
	//db.First(&firstUser)

	t := db.Where("'username' = ?", "Ivan").Find(&User)
	fmt.Println(t.Rows())
	fmt.Println("agaga")
}
