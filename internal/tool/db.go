package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	comment "minitok/cmd/comment/dal"
	favorite "minitok/cmd/favorite/dal"
	user "minitok/cmd/user/dal"
	video "minitok/cmd/video/dal"
	"os"
)

func dsn(username string, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		username,
		os.Getenv("DB_PASSWORD"),
		"localhost",
		3306,
		dbname,
		"charset=utf8&parseTime=True&loc=Local")
}

func main() {
	DB, err := gorm.Open(mysql.Open(dsn("xiayi", "minitok")),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if !m.HasTable(&user.User{}) {
		if err = m.CreateTable(&user.User{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&video.Video{}) {
		if err = m.CreateTable(&video.Video{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&comment.Comment{}) {
		if err = m.CreateTable(&comment.Comment{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&favorite.Favorite{}) {
		if err = m.CreateTable(&favorite.Favorite{}); err != nil {
			panic(err)
		}
	}

}
