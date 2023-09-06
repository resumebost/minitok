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

func GetVideo(vID int64, ctx context.Context) (*Video, error) {
	var video Video
	result := GormDB.WithContext(ctx).First(&video, vID)
	return &video, result.Error
}

func GetVideosByIDs(vIDs []int64, ctx context.Context) ([]*Video, error) {
	var videos []*Video
	result := GormDB.WithContext(ctx).Find(&videos, vIDs)
	return videos, result.Error
}

func GetVideosByAuthorDescByTime(authorID int64, ctx context.Context) ([]*Video, error) {
	var videos []*Video
	result := GormDB.WithContext(ctx).
		Where("author_id = ?", authorID).
		Order("created_at desc").
		Find(&videos)
	return videos, result.Error
}

func GetVideoIdsByAuthor(authorID int64, ctx context.Context) ([]int64, error) {
	var videoIds []int64
	result := GormDB.Model(&Video{}).WithContext(ctx).
		Where("author_id = ?", authorID).
		Pluck("id", &videoIds)
	return videoIds, result.Error
}

func GetVideosDescByTimeLimit(latestTime int64, videoNum int, ctx context.Context) ([]*Video, error) {
	var videos []*Video

	// videoTime := time.Unix(latestTime, 0).Format("2006-01-02 15:04:05")
	result := GormDB.WithContext(ctx).
		Where("unix_timestamp(created_at) < ?", latestTime).
		Order("created_at desc").
		Limit(videoNum).
		Find(&videos)
	return videos, result.Error
}
