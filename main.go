package main

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgtype"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	ID       int    `gorm:"column:Id" gorm:"primaryKey"`
	Username string `gorm:"column:Username"`
	Password string `gorm:"column:Password"`
	Role     int    `gorm:"column:Role"`
}

type Test struct {
	ID       int `gorm:"primaryKey"`
	test1    string
	text2    string
	datetime pgtype.Date
}

type Queue struct {
	ID        int
	Name      string
	TestField string
}

type Sea struct {
	ID   int
	Path string
	Qp   int
	Tt   string
	Pp   int
}

type Inter interface {
}

func (u User) TableName() string {
	return "User"
}

func (u Test) TableName() string {
	return "testTable"
}

func TestHadnler(w http.ResponseWriter, r *http.Request) {

}

func QueueHandler(w http.ResponseWriter, r *http.Request) {

}

func SeaHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//var user1 User
	//var users []User
	//var test Test
	//var inte Inter
	//dsn := "host=localhost user=postgres password=qwerty dbname=postgres port=5432"

	//Init
	r := mux.NewRouter()
	r.HandleFunc("/", TestHadnler)
	r.HandleFunc("/queue", QueueHandler)
	r.HandleFunc("/sea", SeaHandler)
	http.Handle("/", r)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//INSERT
	//us := User{Username: "ustat", Password: "pwd", Role: 1}
	//result := db.Create(&us)
	//fmt.Println(result.RowsAffected)
	//db.Select("Username", "Password", "Role").Create(&us)
	//db.Select('')
	//firstUser := User{}
	//db.First(&firstUser)

	//SELECT
	//result2 := db.Find(&users)
	//fmt.Println(result2.RowsAffected)
	//results := db.Where("'Username' = ?", "Ivan").First(&user1)
	//fmt.Println(results.RowsAffected)
	//
	//resultTest := db.Where("text2 = ?", "2").Find(&inte)
	//fmt.Println(resultTest.RowsAffected)
	//fmt.Println("agaga")
}
