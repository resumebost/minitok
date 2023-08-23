package service

import (
	"context"
	"minitok/cmd/video/dal"
	"minitok/cmd/video/rpc"
	"minitok/internal/unierr"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/favorite"
	"minitok/kitex_gen/user"
	"minitok/kitex_gen/video"
)

type VideoListService struct {
	ctx context.Context
}

func NewVideoListService(ctx context.Context) *VideoListService {
	return &VideoListService{
		ctx: ctx,
	}
}

func (s *VideoListService) GetVideos(req *video.GetVideosRequest) ([]*video.Video, error) {

	videoIds := req.VideoIds

	videos, err := dal.GetVideosByIDs(videoIds, s.ctx)
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return []*video.Video{}, nil
	}
	if len(videos) != len(videoIds) {
		return nil, unierr.VideoNotFound
	}

	res := make([]*video.Video, len(videos))
	//获取视频点赞数
	favoriteCount, err := rpc.CountFavorite(s.ctx,
		&favorite.CountRequest{VideoIdList: videoIds})
	if err != nil {
		return nil, unierr.NewErrCore(
			favoriteCount.StatusCode,
			favoriteCount.StatusMsg)
	}
	//获取视频评论数
	commentCount, err := rpc.CountComment(s.ctx,
		&comment.CountRequest{VideoIdList: videoIds})
	if err != nil {
		return nil, unierr.NewErrCore(
			commentCount.StatusCode,
			commentCount.StatusMsg)
	}

	for i, v := range videos {
		res[i] = &video.Video{
			Id:            int64(v.ID),
			Author:        &user.User{}, //空
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: favoriteCount.FavoriteCountList[i],
			CommentCount:  commentCount.CommentCountList[i],
			IsFavorite:    false, //默认未关注
			Title:         v.Title,
		}
	}
	return res, nil
}
