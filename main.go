package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Username string
	Password string
	Role     int
}

func main() {

	dsn := "host=localhost user=postgres password=qwerty dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	fmt.Println("agaga")
}
