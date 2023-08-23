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

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{
		ctx: ctx,
	}
}

func (s *PublishListService) PublishList(req *video.PublishListRequest) ([]*video.Video, error) {
	userId := req.UserId

	videos, err := dal.GetVideosByAuthorDescByTime(userId, s.ctx)
	if err != nil {
		return nil, err
	}
	//没有投稿过
	if len(videos) == 0 {
		return []*video.Video{}, nil
	}

	videoIds := make([]int64, len(videos))
	for i, v := range videos {
		videoIds[i] = int64(v.ID)
	}
	res := make([]*video.Video, len(videos))

	//获取视频点赞数
	favoriteCount, err := rpc.CountFavorite(s.ctx, &favorite.CountRequest{VideoIdList: videoIds})
	if err != nil {
		return nil, unierr.NewErrCore(
			favoriteCount.StatusCode,
			favoriteCount.StatusMsg)
	}
	//获取视频评论数
	commentCount, err := rpc.CountComment(s.ctx, &comment.CountRequest{VideoIdList: videoIds})
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
	//获取视频作者信息
	author, err := rpc.GetUserInfo(s.ctx, &user.InfoRequest{UserId: userId, Token: req.Token})
	if err != nil {
		return nil, unierr.NewErrCore(
			author.StatusCode,
			author.StatusMsg)
	}

	for i, v := range videos {
		res[i] = &video.Video{
			Id:            int64(v.ID),
			Author:        author.User, //同一用户发布的视频作者相同，都为用户自己
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
