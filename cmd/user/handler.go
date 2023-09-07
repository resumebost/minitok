package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"math/rand"
	"minitok/cmd/user/dal"
	"minitok/cmd/user/pkg/snowflake"
	"minitok/cmd/user/rpc"
	"minitok/cmd/user/tool"
	"minitok/internal/jwt"
	"minitok/internal/unierr"
	"minitok/kitex_gen/favorite"
	user "minitok/kitex_gen/user"
	"minitok/kitex_gen/video"
	"sync"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	var Username = req.Username
	usr, err := dal.GetUserByNameByRegister(ctx, Username)
	if err != nil {
		res := &user.RegisterResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	} else if usr.Username != "" {
		res := &user.RegisterResponse{
			StatusCode: unierr.UsernameExist.ErrCode,
			StatusMsg:  unierr.UsernameExist.ErrMsg,
		}
		return res, nil
	}

	userID := snowflake.GenID()

	// 创建user
	usr = &dal.User{
		ID: userID,
		//ID:       userID,
		Username: req.Username,
		Password: tool.EncryptPassword(req.Password),
		Avatar:   fmt.Sprintf("default%d.png", rand.Intn(10)),
	}

	if err := dal.CreateUser(ctx, usr); err != nil {
		res := &user.RegisterResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	}

	token, err := jwt.GenToken(usr.ID, usr.Username)
	if err != nil {
		res := &user.RegisterResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	}

	res := &user.RegisterResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  unierr.SuccessCode.ErrMsg,
		UserId:     usr.ID,
		Token:      token,
	}
	return res, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {

	// 根据用户名获取密码
	usr, err := dal.GetUserByNameByLogin(ctx, req.Username)
	if err != nil {
		res := &user.LoginResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	} else if usr.Password == "" {
		res := &user.LoginResponse{
			StatusCode: unierr.UsernameNotExist.ErrCode,
			StatusMsg:  unierr.UsernameNotExist.ErrMsg,
		}
		return res, nil
	}

	if tool.EncryptPassword(req.Password) != usr.Password {
		res := &user.LoginResponse{
			StatusCode: unierr.PasswordWrong.ErrCode,
			StatusMsg:  unierr.PasswordWrong.ErrMsg,
		}
		return res, nil
	}

	usr = &dal.User{
		ID:       usr.ID,
		Username: req.Username,
	}

	token, err := jwt.GenToken(usr.ID, usr.Username)
	if err != nil {
		res := &user.LoginResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	}

	res := &user.LoginResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  unierr.SuccessCode.ErrMsg,
		UserId:     usr.ID,
		Token:      token,
	}
	return res, nil

}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	var token = req.Token

	claims, err := jwt.ParseToken(token)
	if err != nil {
		return
	}
	userID := claims.ID

	usr, err := dal.GetUserByID(ctx, userID)
	if err != nil {
		res := &user.InfoResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	} else if usr.ID == 0 {
		res := &user.InfoResponse{
			StatusCode: unierr.UserNotExist.ErrCode,
			StatusMsg:  unierr.UserNotExist.ErrMsg,
			User:       nil,
		}
		return res, nil
	}
	rUser := user.User{
		Id:              usr.ID,
		Name:            usr.Username,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          usr.Avatar,
		BackgroundImage: usr.BackgroundImage,
		Signature:       usr.Signature,
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	//视频数量和获赞数量
	go func() {
		defer wg.Done()
		//视频数量
		reqVideo := &video.PublishListIdsRequest{UserId: req.UserId}
		respPublishListIds, err := rpc.GetPublishListIds(ctx, reqVideo)
		if err != nil {
			resp = &user.InfoResponse{
				StatusCode: unierr.SuccessCode.ErrCode,
				StatusMsg:  unierr.SuccessCode.ErrMsg,
				User:       nil,
			}
			return
		}
		rUser.WorkCount = int64(len(respPublishListIds.VideoIdList))

		//获赞数量
		reqFavoriteCount := &favorite.CountRequest{VideoIdList: respPublishListIds.VideoIdList}
		respFavoriteCount, err := rpc.FavoriteCount(ctx, reqFavoriteCount)
		if err != nil {
			resp = &user.InfoResponse{
				StatusCode: unierr.SuccessCode.ErrCode,
				StatusMsg:  unierr.SuccessCode.ErrMsg,
				User:       nil,
			}
			return
		}

		var sum int64 = 0
		for _, c := range respFavoriteCount.FavoriteCountList {
			sum += c
		}
		rUser.TotalFavorited = sum
	}()

	//点赞数量
	go func() {
		defer wg.Done()
		reqCountByUser := &favorite.CountByUserRequest{UserId: req.UserId}
		respCountByUser, err := rpc.FavoriteCountByUser(ctx, reqCountByUser)
		if err != nil {
			resp = &user.InfoResponse{
				StatusCode: unierr.SuccessCode.ErrCode,
				StatusMsg:  unierr.SuccessCode.ErrMsg,
				User:       nil,
			}
			return
		}
		rUser.FavoriteCount = respCountByUser.FavoriteCount

	}()
	wg.Wait()

	res := &user.InfoResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  unierr.SuccessCode.ErrMsg,
		User:       &rUser,
	}
	return res, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	userID := req.UserId

	usr, err := dal.GetUserByID(ctx, userID)
	if err != nil {
		res := &user.GetUserResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	} else if usr.ID == 0 {
		res := &user.GetUserResponse{
			StatusCode: unierr.UserNotExist.ErrCode,
			StatusMsg:  unierr.UserNotExist.ErrMsg,
			User:       nil,
		}
		return res, nil
	}

	rUser := user.User{
		Id:              usr.ID,
		Name:            usr.Username,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          usr.Avatar,
		BackgroundImage: usr.BackgroundImage,
		Signature:       usr.Signature,
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	//视频数量和获赞数量
	go func() {
		defer wg.Done()
		//视频数量
		respPublishListIds, err := rpc.GetPublishListIds(ctx,
			&video.PublishListIdsRequest{UserId: req.UserId})
		if err != nil {
			resp = &user.GetUserResponse{
				StatusCode: unierr.SuccessCode.ErrCode,
				StatusMsg:  unierr.SuccessCode.ErrMsg,
				User:       nil,
			}
			return
		}
		rUser.WorkCount = int64(len(respPublishListIds.VideoIdList))

		//获赞数量
		respFavoriteCount, err := rpc.FavoriteCount(ctx,
			&favorite.CountRequest{VideoIdList: respPublishListIds.VideoIdList})
		if err != nil {
			resp = &user.GetUserResponse{
				StatusCode: unierr.SuccessCode.ErrCode,
				StatusMsg:  unierr.SuccessCode.ErrMsg,
				User:       nil,
			}
			return
		}

		var sum int64 = 0
		for _, c := range respFavoriteCount.FavoriteCountList {
			sum += c
		}
		rUser.TotalFavorited = sum
	}()

	//点赞数量
	go func() {
		defer wg.Done()
		respCountByUser, err := rpc.FavoriteCountByUser(ctx,
			&favorite.CountByUserRequest{UserId: req.UserId})
		if err != nil {
			resp = &user.GetUserResponse{
				StatusCode: unierr.SuccessCode.ErrCode,
				StatusMsg:  unierr.SuccessCode.ErrMsg,
				User:       nil,
			}
			return
		}
		rUser.FavoriteCount = respCountByUser.FavoriteCount

	}()
	wg.Wait()

	res := &user.GetUserResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  unierr.SuccessCode.ErrMsg,
		User:       &rUser,
	}
	return res, nil
}

// GetUsers implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUsers(ctx context.Context, req *user.GetUsersRequest) (resp *user.GetUsersResponse, err error) {
	userIDList := req.UserIdList

	usrList, err := dal.GetUsersByIDList(ctx, userIDList)
	if err != nil {
		res := &user.GetUsersResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
		return res, nil
	} else if len(usrList) != len(userIDList) {
		fmt.Println(usrList)
		fmt.Println(userIDList)
		res := &user.GetUsersResponse{
			StatusCode: unierr.UserNotExist.ErrCode,
			StatusMsg:  unierr.UserNotExist.ErrMsg,
		}
		return res, nil
	}

	var workCountList = make([]int64, len(userIDList))
	var totalFavoriteList = make([]int64, len(userIDList))
	var favoriteCountList = make([]int64, len(userIDList))
	var rUsers = make([]*user.User, len(userIDList))

	wg := sync.WaitGroup{}
	wg.Add(2)

	//视频数量和获赞数量
	go func() {
		defer wg.Done()
		//视频数量
		for i, u := range userIDList {
			respPublishListIds, err := rpc.GetPublishListIds(ctx,
				&video.PublishListIdsRequest{UserId: u})
			if err != nil {
				klog.Fatalf("get PublishListIds failed: %s", err)
				return
			}
			workCountList[i] = int64(len(respPublishListIds.VideoIdList))

			//获赞数量
			respFavoriteCount, err := rpc.FavoriteCount(ctx,
				&favorite.CountRequest{VideoIdList: respPublishListIds.VideoIdList})
			if err != nil {
				klog.Fatalf("get FavoriteCount failed: %s", err)
				return
			}
			var sum int64 = 0
			for _, c := range respFavoriteCount.FavoriteCountList {
				sum += c
			}
			totalFavoriteList[i] = sum
		}

	}()

	//点赞数量
	go func() {
		defer wg.Done()
		for i, u := range userIDList {
			respCountByUser, err := rpc.FavoriteCountByUser(ctx,
				&favorite.CountByUserRequest{UserId: u})
			if err != nil {
				klog.Fatal("get CountByUser failed: %s", err)
				return
			}
			favoriteCountList[i] = respCountByUser.FavoriteCount
		}
	}()
	wg.Wait()

	for i, usr := range usrList {
		rUsers[i] = &user.User{
			Id:              usr.ID,
			Name:            usr.Username,
			FollowCount:     0,
			FollowerCount:   0,
			IsFollow:        false,
			Avatar:          usr.Avatar,
			BackgroundImage: usr.BackgroundImage,
			Signature:       usr.Signature,
			TotalFavorited:  totalFavoriteList[i],
			WorkCount:       workCountList[i],
			FavoriteCount:   favoriteCountList[i],
		}
	}

	res := &user.GetUsersResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  unierr.SuccessCode.ErrMsg,
		User:       rUsers,
	}
	return res, nil
}
