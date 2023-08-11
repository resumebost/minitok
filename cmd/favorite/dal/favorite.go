package dal

import (
	"gorm.io/gorm"
	"minitok/internal/dal"
)

var GormDB *gorm.DB

type Favorite struct {
	gorm.Model

	VideoID string
	UserID  string
}

func SetFavoriteDB() {
	GormDB = dal.InitGorm()
}
