package dal

import (
	"gorm.io/gorm"
	"minitok/internal/dal"
)

var GormDB *gorm.DB

type User struct {
	gorm.Model

	Username        string
	Password        string
	Avatar          string
	BackgroundImage string
	Signature       string
}

func SetUserDB() {
	GormDB = dal.InitGorm()
}
