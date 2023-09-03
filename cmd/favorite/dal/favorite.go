package dal

import (
	"context"
	"minitok/internal/dal"

	"gorm.io/gorm"
)

var GormDB *gorm.DB

type Favorite struct {
	gorm.Model

	VideoID int64 `json:"video_id"`
	UserID  int64 `json:"user_id"`
}

func SetFavoriteDB() {
	GormDB = dal.InitGorm()
}

func CreateFavorite(ctx context.Context, favorites []*Favorite) error {
	// db := GormDB.WithContext(ctx) //方便链路追踪
	if err := GormDB.WithContext(ctx).Create(favorites).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFavorite(ctx context.Context, userID int64, videoID int64) error {
	db := GormDB.WithContext(ctx)
	return db.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&Favorite{}).Error
}

// GetUserLikedVideoIDs retrieves video IDs that the user has liked.
func GetUserLikedVideoIDs(ctx context.Context, userID int64) ([]int64, error) {
	var likedVideoIDs []int64

	db := GormDB.WithContext(ctx)
	err := db.Model(&Favorite{}).Select("video_id").Where("user_id = ?", userID).Find(&likedVideoIDs).Error

	if err != nil {
		return nil, err
	}

	return likedVideoIDs, nil
}

// 列表判断
// JudgeLikes checks whether the user has liked the given videos.
func JudgeLikes(ctx context.Context, userID int64, videoIDs []int64) ([]bool, error) {
	var results []bool

	db := GormDB.WithContext(ctx)

	for _, videoID := range videoIDs {
		var count int64
		err := db.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userID, videoID).Count(&count).Error
		if err != nil {
			return nil, err
		}
		results = append(results, count > 0)
	}

	return results, nil
}

// 列表总数
// GetFavoriteCounts retrieves the favorite counts for the given videos.
func GetFavoriteCounts(ctx context.Context, videoIDs []int64) ([]int64, error) {
	db := GormDB.WithContext(ctx)

	// 用 map 来记录视频ID和对应的点赞数
	videoIDToCount := make(map[int64]int64)

	rows, err := db.Model(&Favorite{}).
		Select("video_id, COUNT(*) as favorite_count").
		Where("video_id IN (?)", videoIDs).
		Group("video_id").
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 将videoID遍历并更新对应的map条目
	for rows.Next() {
		var videoID, count int64
		if err := rows.Scan(&videoID, &count); err == nil {
			videoIDToCount[videoID] = count
		}
	}

	// 创建结果列表，如果视频ID在map中不存在，则返回0
	var counts []int64
	for _, videoID := range videoIDs {
		count, found := videoIDToCount[videoID]
		if !found {
			count = 0
		}
		counts = append(counts, count)
	}

	return counts, nil
}

func CountFavoritesByUser(ctx context.Context, userID int64) (int64, error) {
	db := GormDB.WithContext(ctx)

	var count int64
	if err := db.Model(&Favorite{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
