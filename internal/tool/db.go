package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "minitok/cmd/comment/dal"
	. "minitok/cmd/favorite/dal"
	. "minitok/cmd/user/dal"
	. "minitok/cmd/video/dal"
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
	if !m.HasTable(&User{}) {
		if err = m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&Video{}) {
		if err = m.CreateTable(&Video{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&Comment{}) {
		if err = m.CreateTable(&Comment{}); err != nil {
			panic(err)
		}
	}

	if !m.HasTable(&Favorite{}) {
		if err = m.CreateTable(&Favorite{}); err != nil {
			panic(err)
		}
	}

}
