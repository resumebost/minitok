package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"minitok/cmd/video/dal"
	"minitok/cmd/video/rpc"
	"minitok/internal/unierr"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/favorite"
	"minitok/kitex_gen/user"
	"minitok/kitex_gen/video"
	"sync"
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

	authorIds := make([]int64, len(videoIds))
	for i, v := range videos {
		authorIds[i] = v.AuthorID
	}

	res := make([]*video.Video, len(videos))
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
		isFavorite, err = rpc.JudgeFavorite(s.ctx, &favorite.JudgeRequest{
			VideoIdList: videoIds,
			Token:       req.Token})
		if err != nil {
			klog.Fatalf("judge favorite failed: %s", err)
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
	return res, nil
}
