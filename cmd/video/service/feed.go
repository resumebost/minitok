package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"minitok/cmd/video/dal"
	"minitok/cmd/video/rpc"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/favorite"
	"minitok/kitex_gen/user"
	"minitok/kitex_gen/video"
	"sync"
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
	token := req.Token

	videos, err := dal.GetVideosDescByTimeLimit(latestTime, FEED_LIMIT, s.ctx)
	if err != nil {
		return nil, 0, err
	}

	//无视频，返回next_time为原时间戳
	videoNum := len(videos)
	if videoNum == 0 {
		return []*video.Video{}, latestTime, nil
	}

	videoIds := make([]int64, videoNum)
	authorIds := make([]int64, videoNum)
	for i, v := range videos {
		videoIds[i] = int64(v.ID)
		authorIds[i] = v.AuthorID
	}
	res := make([]*video.Video, videoNum)

	nextTime := videos[videoNum-1].CreatedAt.UTC().UnixMilli()

	var favoriteCount *favorite.CountResponse
	var commentCount *comment.CountResponse
	var isFavorite *favorite.JudgeResponse
	var authors *user.GetUsersResponse

	var wg sync.WaitGroup
	wg.Add(4)
	//获取视频点赞数
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("count favorite failed: %s", err)
			}
		}()
		defer wg.Done()
		favoriteCount, err = rpc.CountFavorite(s.ctx, &favorite.CountRequest{VideoIdList: videoIds})
		if err != nil {
			klog.Fatalf("count favorite failed: %s", err)
		}
	}()

	//获取视频评论数
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("count comment failed: %s", err)
			}
		}()
		defer wg.Done()
		commentCount, err = rpc.CountComment(s.ctx, &comment.CountRequest{VideoIdList: videoIds})
		if err != nil {
			klog.Fatalf("count comment failed: %s", err)
		}
	}()

	//获取用户是否点赞过视频
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("judge favorite failed: %s", err)
			}
		}()
		defer wg.Done()
		if len(token) == 0 {
			isLikeList := make([]bool, videoNum)
			isFavorite = &favorite.JudgeResponse{Is_LikeList: isLikeList}
		} else {
			isFavorite, err = rpc.JudgeFavorite(s.ctx, &favorite.JudgeRequest{
				VideoIdList: videoIds,
				Token:       token})
			if err != nil {
				klog.Fatalf("judge favorite failed1: %s", err)
			}
		}
	}()

	//获取视频作者信息
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("get authors failed: %s", err)
			}
		}()
		defer wg.Done()
		authors, err = rpc.GetUsers(s.ctx, &user.GetUsersRequest{
			UserIdList: authorIds})
		if err != nil {
			klog.Fatalf("get authors failed1: %s", err)
		}
	}()

	wg.Wait()

	for i, v := range videos {

		res[i] = &video.Video{
			Id:     int64(v.ID),
			Author: authors.User[i],
			//Author:        &user.User{},
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
