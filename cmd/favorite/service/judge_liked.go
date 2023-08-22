package service

import (
	"context"
	"minitok/cmd/favorite/dal"
)

type JudgeLikedVideosService struct {
	ctx context.Context
}

// 初始化ctx
func NewJudgeLikedVideosService(ctx context.Context) *JudgeLikedVideosService {
	return &JudgeLikedVideosService{ctx: ctx}
}

//判断点赞与否
func (s *JudgeLikedVideosService)JudgeLikedVideos( userID int64, videoIDs []int64) ([]bool, error) {
	isLikedList, err := dal.JudgeLikes(s.ctx, userID, videoIDs)
	if err != nil {
		return nil, err
	}
	return isLikedList, nil
}

//视频点赞总数
func (s *JudgeLikedVideosService)GetVideoFavoriteCounts(videoIDs []int64) ([]int64, error) {
	favoriteCounts, err := dal.GetFavoriteCounts(s.ctx, videoIDs)
	if err != nil {
		return nil, err
	}
	return favoriteCounts, nil
}
