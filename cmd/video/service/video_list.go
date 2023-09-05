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
	//获取用户是否点赞过视频
	isFavorite, err := rpc.JudgeFavorite(s.ctx, &favorite.JudgeRequest{
		VideoIdList: videoIds,
		Token:       req.Token})
	if err != nil {
		return nil, unierr.NewErrCore(
			isFavorite.StatusCode,
			isFavorite.StatusMsg)
	}

	for i, v := range videos {
		author, err := rpc.GetUserInfo(s.ctx, &user.InfoRequest{
			UserId: v.AuthorID,
			Token:  req.Token})
		if err != nil {
			return nil, unierr.NewErrCore(
				author.StatusCode,
				author.StatusMsg)
		}

		res[i] = &video.Video{
			Id:     int64(v.ID),
			Author: author.User,
			//Author:        &user.User{},
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: favoriteCount.FavoriteCountList[i],
			CommentCount:  commentCount.CommentCountList[i],
			IsFavorite:    isFavorite.Is_LikeList[i],
			Title:         v.Title,
		}
	}
	return res, nil
}
