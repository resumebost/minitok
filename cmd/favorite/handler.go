package main

import (
	"context"
	"errors"
	"minitok/cmd/favorite/service"
	"minitok/internal/jwt"
	"minitok/internal/unierr"
	favorite "minitok/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// 点赞&取消
// Action implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Action(ctx context.Context, req *favorite.ActionRequest) (resp *favorite.ActionResponse, err error) {

	// userID := ctx.Value("id").(int64) // 用户id
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		return
	}
	userID := claims.ID

	resp = new(favorite.ActionResponse)
	if req.ActionType == 1 { //  1 represents "like" action
		err = service.NewLikeVideoService(ctx).LikeVideo(userID, req.VideoId)
	} else if req.ActionType == 2 { //  2 represents "unlike" action
		err = service.NewLikeVideoService(ctx).UnlikeVideo(userID, req.VideoId)
	} else {
		err = errors.New("invalid action type , input 1 or 2")
		resp = &favorite.ActionResponse{
			StatusCode: unierr.IllegalParams.ErrCode,
			StatusMsg:  unierr.IllegalParams.ErrMsg,
		}
		return resp, err
	}

	if err != nil {
		resp = &favorite.ActionResponse{
			StatusCode: unierr.FavoriteActionError.ErrCode,
			StatusMsg:  unierr.FavoriteActionError.ErrMsg,
		}
		return resp, err
	}

	// resp数据封装
	resp = &favorite.ActionResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  "Action successful",
	}
	return resp, nil
}

// 点赞列表
// List implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) List(ctx context.Context, req *favorite.ListRequest) (resp *favorite.ListResponse, err error) {
	// 获取需要查看的用户信息
	userID := req.UserId

	// 调用 service 获取喜欢的视频列表
	likedVideos, err := service.NewLikeVideoListService(ctx).GetLikedVideos(req.Token, userID)
	if err != nil {
		// 处理错误并返回错误响应
		return &favorite.ListResponse{
			StatusCode: unierr.GetFavoriteVideoListFiled.ErrCode,
			StatusMsg:  unierr.GetFavoriteVideoListFiled.ErrMsg,
		}, err
	}

	// 返回响应
	return &favorite.ListResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  "Action successful",
		VideoList:  likedVideos,
	}, nil
}

// 点赞与否
// Judge implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Judge(ctx context.Context, req *favorite.JudgeRequest) (resp *favorite.JudgeResponse, err error) {
	// userID := ctx.Value("id").(int64) // 用户id
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		return
	}
	userID := claims.ID

	videoIDs := req.VideoIdList

	isLikedList, err := service.NewJudgeLikedVideosService(ctx).JudgeLikedVideos(userID, videoIDs)
	if err != nil {
		return &favorite.JudgeResponse{
			StatusCode:  unierr.GetFavoriteJudgeFiled.ErrCode,
			StatusMsg:   unierr.GetFavoriteJudgeFiled.ErrMsg,
			Is_LikeList: nil, // Handle error response
		}, err
	}

	return &favorite.JudgeResponse{
		StatusCode:  unierr.SuccessCode.ErrCode,
		StatusMsg:   "Action successful",
		Is_LikeList: isLikedList,
	}, nil
}

// 点赞总数
// Count implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Count(ctx context.Context, req *favorite.CountRequest) (resp *favorite.CountResponse, err error) {

	videoIDs := req.VideoIdList

	favoriteCounts, err := service.NewJudgeLikedVideosService(ctx).GetVideoFavoriteCounts(videoIDs)
	if err != nil {
		return &favorite.CountResponse{
			StatusCode:        unierr.GetFavoriteVideoCountFiled.ErrCode,
			StatusMsg:         unierr.GetFavoriteVideoCountFiled.ErrMsg,
			FavoriteCountList: nil, // Handle error response
		}, err
	}

	return &favorite.CountResponse{
		StatusCode:        unierr.SuccessCode.ErrCode,
		StatusMsg:         "Action successful",
		FavoriteCountList: favoriteCounts,
	}, nil
}

// CountByUser implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CountByUser(ctx context.Context, req *favorite.CountByUserRequest) (*favorite.CountByUserResponse, error) {
	userID := req.UserId

	favoriteCount, err := service.NewCountByUserService(ctx).CountByUser(userID)
	if err != nil {
		return &favorite.CountByUserResponse{
			StatusCode: unierr.GetUserFavoriteVideoCountFiled.ErrCode,
			StatusMsg:  unierr.GetUserFavoriteVideoCountFiled.ErrMsg,
			FavoriteCount: 0,
		}, err
	}

	return &favorite.CountByUserResponse{
		StatusCode:    unierr.SuccessCode.ErrCode,
		StatusMsg:     "Action successful",
		FavoriteCount: favoriteCount,//视频收藏的总数
	}, nil
}
