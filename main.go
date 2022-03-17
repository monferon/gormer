package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gormer/controllers"
	"gormer/database"

	docs "gormer/docs"
)

//type User struct {
//	ID       int    `gorm:"column:Id" gorm:"primaryKey"`
//	Username string `gorm:"column:Username"`
//	Password string `gorm:"column:Password"`
//	Role     int    `gorm:"column:Role"`
//}
//
//type Test struct {
//	ID       int `gorm:"primaryKey"`
//	test1    string
//	text2    string
//	datetime pgtype.Date
//}

//type Queue struct {
//	ID        int
//	Name      string
//	TestField string
//}
//
//type Sea struct {
//	ID   int
//	Path string
//	Qp   int
//	Tt   string
//	Pp   int
//}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	database.Connect()

	v1 := r.Group("/v1")
	{
		v1.GET("/books", controllers.FindBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.GET("/books/:id", controllers.FindBook)
		v1.PATCH("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
