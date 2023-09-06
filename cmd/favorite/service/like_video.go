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

func (s *LikeVideoService) LikeOrUnlikeVideo(userID int64, videoID int64, actionType int32) error {
	
	// 首先检查用户是否已经点赞了视频
	liked, err := dal.CheckIfUserLikedVideo(s.ctx, userID, videoID)
	if err != nil {
		return err
	}

	// 如果用户已经点赞并且想取消点赞，或者用户未点赞并且想点赞
	if (liked && actionType == 2) || (!liked && actionType == 1) {
		// 用户已经点赞且想取消点赞，或者用户未点赞且想点赞，执行相应操作
		if actionType == 1 {
			// 点赞
			favoriteModel := &dal.Favorite{
				UserID:  userID,
				VideoID: videoID,
			}
			return dal.CreateFavorite(s.ctx, []*dal.Favorite{favoriteModel})
		} else {
			// 取消点赞
			return dal.DeleteFavorite(s.ctx, userID, videoID)
		}
	}

	return nil
}