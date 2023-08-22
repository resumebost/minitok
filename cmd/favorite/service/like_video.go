package service

import (
	"context"
	"minitok/cmd/favorite/dal"
	"minitok/cmd/favorite/pack"
	"minitok/cmd/favorite/rpc"
	"minitok/kitex_gen/video"
)

type LikeVideoListService struct {
	ctx context.Context
}

//初始化
func NewLikeVideoListService(ctx context.Context) *LikeVideoListService {
	return &LikeVideoListService{ctx: ctx}
}

// GetUserLikedVideos retrieves videos that the user has liked.
func (s *LikeVideoListService)GetLikedVideos(userID int64) ([]*video.Video, error) {
	likedVideoIDs, err := dal.GetUserLikedVideoIDs(s.ctx , userID)
	if err != nil {
		return nil, err
	}

	videoMap, err := rpc.GetVideosInfo(s.ctx, likedVideoIDs)
	if err != nil {
		return nil, err
	}

	//打个包
	likedVideos := pack.ConvertToFavoriteVideos(videoMap, likedVideoIDs)
	return likedVideos, nil
}