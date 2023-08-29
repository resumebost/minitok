package main

import (
	"context"
	"fmt"
	"math/rand"
	"minitok/cmd/user/dal"
	"minitok/cmd/user/pkg/snowflake"
	"minitok/cmd/user/tool"
	"minitok/internal/jwt"
	"minitok/internal/unierr"
	user "minitok/kitex_gen/user"
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
	} else if usr.UserName != "" {
		res := &user.RegisterResponse{
			StatusCode: unierr.UsernameExist.ErrCode,
			StatusMsg:  unierr.UsernameExist.ErrMsg,
		}
		return res, nil
	}

	userID := snowflake.GenID()

	// 创建user
	usr = &dal.User{
		UserID:   userID,
		UserName: req.Username,
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

	token, err := jwt.GenToken(usr.UserID, usr.UserName)
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
		UserId:     usr.UserID,
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
		UserID:   usr.UserID,
		UserName: req.Username,
	}

	token, err := jwt.GenToken(usr.UserID, usr.UserName)
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
		UserId:     usr.UserID,
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

	res := &user.InfoResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  unierr.SuccessCode.ErrMsg,
		User: &user.User{
			Id:              usr.UserID,
			Name:            usr.UserName,
			FollowCount:     int64(usr.FollowerCount),
			FollowerCount:   int64(usr.FollowingCount),
			IsFollow:        usr.UserID == userID,
			Avatar:          usr.Avatar,
			BackgroundImage: usr.BackgroundImage,
			Signature:       usr.Signature,
			TotalFavorited:  int64(usr.TotalFavorited),
			WorkCount:       int64(usr.WorkCount),
			FavoriteCount:   int64(usr.FavoriteCount),
		},
	}

	return res, nil
}
