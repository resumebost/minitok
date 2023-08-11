package dal

import (
	"gorm.io/gorm"
	"minitok/internal/dal"
)

var GormDB *gorm.DB

type Comment struct {
	gorm.Model

	UserID  string
	VideoID string
	Content string
}

func SetCommentDB() {
	GormDB = dal.InitGorm()
}
