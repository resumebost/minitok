package service

import (
	"context"
	"minitok/cmd/favorite/dal"
)

type LikeVideoService struct {
	ctx context.Context
}

// 初始化
func NewLikeVideoService(ctx context.Context) *LikeVideoService {
	return &LikeVideoService{ctx: ctx}
}

// 插入数据：点赞
func (s *LikeVideoService) LikeVideo(userID int64, videoID int64) error {
	//将数据封装一下好一点，dal那边使用一次插入多条，但怎么用都ok
	favoriteModel := &dal.Favorite{
		UserID:  userID,
		VideoID: videoID,
	}
	return dal.CreateFavorite(s.ctx, []*dal.Favorite{favoriteModel})
}

func (s *LikeVideoService) UnlikeVideo(userID int64, videoID int64) error {
	return dal.DeleteFavorite(s.ctx, userID, videoID)
}
