package dal

import (
	"context"
	"gorm.io/gorm"
	"minitok/internal/dal"
)

var GormDB *gorm.DB

type Comment struct {
	gorm.Model

	UserID  int64
	VideoID int64
	Content string
}

func SetCommentDB() {
	GormDB = dal.InitGorm()
}

func CreateComment(comment *Comment, ctx context.Context) (*Comment, error) {
	result := GormDB.WithContext(ctx).Create(comment)
	return comment, result.Error
}

func DeleteComment(commentId int64, ctx context.Context) error {
	result := GormDB.WithContext(ctx).Where("id = ?", commentId).Delete(&Comment{})
	return result.Error
}

func CommentList(videoId int64, ctx context.Context) ([]*Comment, error) {
	var comments []*Comment
	result := GormDB.WithContext(ctx).
		Where("video_id = ?", videoId).
		Order("created_at desc").
		Find(&comments)
	return comments, result.Error
}

func CountComments(videoIds []int64, ctx context.Context) ([]int64, error) {
	db := GormDB.WithContext(ctx)

	videoIDToCount := make(map[int64]int64)

	rows, err := db.Model(&Comment{}).
		Select("video_id, COUNT(*) as comment_count").
		Where("video_id IN (?)", videoIds).
		Group("video_id").
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var videoID, count int64
		if err := rows.Scan(&videoID, &count); err == nil {
			videoIDToCount[videoID] = count
		}
	}

	var counts []int64
	for _, videoID := range videoIds {
		count, found := videoIDToCount[videoID]
		if !found {
			count = 0
		}
		counts = append(counts, count)
	}

	return counts, nil
}
