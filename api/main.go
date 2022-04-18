package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var db *sqlx.DB

func init() {
	var err error
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")
	db, err = sqlx.Open("mysql", user+":"+pass+"@tcp("+host+")/"+database+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.GET("/post/list", handlerGetPostList)
	e.GET("/post/:post_id", handlerGetPost)
	e.GET("/post/:post_id/comments", handlerGetComments)

	e.POST("/post", handlerCreatePost)
	e.POST("/post/:post_id/comment", handlerCreateComment)

	e.Logger.Fatal(e.Start(":80"))
}
