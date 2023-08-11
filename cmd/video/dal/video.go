package dal

import (
	"gorm.io/gorm"
	"minitok/internal/dal"
)

var GormDB *gorm.DB

type Video struct {
	gorm.Model

	AuthorID string `gorm:"index"`
	PlayURL  string
	CoverURL string
	Title    string
}

func SetVideoDB() {
	GormDB = dal.InitGorm()
}
