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

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

const FEED_LIMIT = 30

func (s *FeedService) Feed(req *video.FeedRequest) ([]*video.Video, int64, error) {

	latestTime := req.LatestTime

	videos, err := dal.GetVideosDescByTimeLimit(latestTime, FEED_LIMIT, s.ctx)
	if err != nil {
		return nil, 0, err
	}

	//无视频，返回next_time为原时间戳
	videoNum := len(videos)
	if videoNum == 0 {
		return []*video.Video{}, latestTime, nil
	}

	videoIds := make([]int64, len(videos))
	authorIds := make([]int64, len(videos))
	for i, v := range videos {
		videoIds[i] = int64(v.ID)
		authorIds[i] = v.AuthorID
	}
	res := make([]*video.Video, len(videos))
	nextTime := videos[videoNum-1].CreatedAt.Unix()

	//获取视频点赞数
	favoriteCount, err := rpc.CountFavorite(s.ctx, &favorite.CountRequest{VideoIdList: videoIds})
	if err != nil {
		return nil, 0, unierr.NewErrCore(
			favoriteCount.StatusCode,
			favoriteCount.StatusMsg)
	}
	//获取视频评论数
	commentCount, err := rpc.CountComment(s.ctx, &comment.CountRequest{VideoIdList: videoIds})
	if err != nil {
		return nil, 0, unierr.NewErrCore(
			commentCount.StatusCode,
			commentCount.StatusMsg)
	}
	//获取用户是否点赞过视频
	isFavorite, err := rpc.JudgeFavorite(s.ctx, &favorite.JudgeRequest{
		VideoIdList: videoIds,
		Token:       req.Token})
	if err != nil {
		return nil, 0, unierr.NewErrCore(
			isFavorite.StatusCode,
			isFavorite.StatusMsg)
	}

	for i, v := range videos {
		//获取视频作者信息
		author, err := rpc.GetUserInfo(s.ctx, &user.InfoRequest{
			UserId: v.AuthorID,
			Token:  req.Token})
		if err != nil {
			return nil, 0, unierr.NewErrCore(
				author.StatusCode,
				author.StatusMsg)
		}

		res[i] = &video.Video{
			Id:            int64(v.ID),
			Author:        author.User,
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: favoriteCount.FavoriteCountList[i],
			CommentCount:  commentCount.CommentCountList[i],
			IsFavorite:    isFavorite.Is_LikeList[i],
			Title:         v.Title,
		}
	}
	return res, nextTime, nil
}
