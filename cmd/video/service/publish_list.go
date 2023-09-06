package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"minitok/cmd/video/dal"
	"minitok/cmd/video/rpc"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/favorite"
	"minitok/kitex_gen/user"
	"minitok/kitex_gen/video"
	"sync"
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
	fmt.Println(videos)
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

	var favoriteCount *favorite.CountResponse
	var commentCount *comment.CountResponse
	var isFavorite *favorite.JudgeResponse
	var author *user.InfoResponse

	var wg sync.WaitGroup
	wg.Add(4)
	//获取视频点赞数
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("count favorite filed: %s", err)
			}
		}()
		defer wg.Done()
		favoriteCount, err = rpc.CountFavorite(s.ctx, &favorite.CountRequest{VideoIdList: videoIds})
		if err != nil {
			klog.Fatalf("count favorite filed: %s", err)
		}
	}()

	//获取视频评论数
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("count comment filed: %s", err)
			}
		}()
		defer wg.Done()
		commentCount, err = rpc.CountComment(s.ctx, &comment.CountRequest{VideoIdList: videoIds})
		if err != nil {
			klog.Fatalf("count comment filed: %s", err)
		}
	}()

	//获取用户是否点赞过视频
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("judge favorite filed: %s", err)
			}
		}()
		defer wg.Done()
		isFavorite, err = rpc.JudgeFavorite(s.ctx, &favorite.JudgeRequest{
			VideoIdList: videoIds,
			Token:       req.Token})
		if err != nil {
			klog.Fatalf("judge favorite filed: %s", err)
		}
	}()

	//获取视频作者信息
	go func() {
		defer func() {
			if err := recover(); err != nil {
				klog.Fatalf("get userInfo filed: %s", err)
			}
		}()
		defer wg.Done()
		author, err = rpc.GetUserInfo(s.ctx, &user.InfoRequest{UserId: userId, Token: req.Token})
		if err != nil {
			klog.Fatalf("get userInfo filed: %s", err)
		}
	}()

	wg.Wait()

	for i, v := range videos {
		res[i] = &video.Video{
			Id:     int64(v.ID),
			Author: author.User, //同一用户发布的视频作者相同，都为用户自己
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

func (s *PublishListService) PublishListIds(req *video.PublishListIdsRequest) ([]int64, error) {
	authorId := req.UserId

	videoIdList, err := dal.GetVideoIdsByAuthor(authorId, s.ctx)
	if err != nil {
		return nil, err
	}

	return videoIdList, nil
}
