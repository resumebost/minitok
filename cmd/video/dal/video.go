package dal

import (
	"context"
	"gorm.io/gorm"
	"minitok/internal/dal"
)

var GormDB *gorm.DB

type Video struct {
	gorm.Model

	AuthorID int64 `gorm:"index"`
	PlayURL  string
	CoverURL string
	Title    string
}

func SetVideoDB() {
	GormDB = dal.InitGorm()
}

func CreateVideo(video *Video, ctx context.Context) (int64, error) {
	result := GormDB.WithContext(ctx).Create(video)
	return int64(video.ID), result.Error
}
