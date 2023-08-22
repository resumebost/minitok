package main

import (
	"context"
	"minitok/cmd/video/service"
	"minitok/internal/unierr"
	video "minitok/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = nil

	videoList, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp = &video.FeedResponse{StatusCode: unierr.GetFeedFiled.ErrCode, StatusMsg: unierr.GetFeedFiled.ErrMsg}
		return resp, err
	}

	resp = &video.FeedResponse{StatusCode: unierr.SuccessCode.ErrCode, StatusMsg: unierr.SuccessCode.ErrMsg, VideoList: videoList, NextTime: nextTime}
	return resp, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	//fmt.Println("进入handler")
	resp = nil

	err = service.NewUploadVideoService(ctx).PublishVideo(req)
	if err != nil {
		resp = &video.PublishActionResponse{StatusCode: unierr.VideoPublishFiled.ErrCode, StatusMsg: unierr.VideoPublishFiled.ErrMsg}
		//fmt.Println("videp handler: " + err.Error())
		return resp, err
	}

	resp = &video.PublishActionResponse{StatusCode: unierr.SuccessCode.ErrCode, StatusMsg: unierr.SuccessCode.ErrMsg}
	//fmt.Println("video服务无问题")
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	resp = nil

	videoList, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp = &video.PublishListResponse{StatusCode: unierr.GetVideoListFiled.ErrCode, StatusMsg: unierr.GetVideoListFiled.ErrMsg}
		return resp, err
	}

	resp = &video.PublishListResponse{StatusCode: unierr.SuccessCode.ErrCode, StatusMsg: unierr.SuccessCode.ErrMsg, VideoList: videoList}
	return resp, nil
}

// GetVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideos(ctx context.Context, req *video.GetVideosRequest) (resp *video.GetVideosResponse, err error) {
	resp = nil

	videoList, err := service.NewVideoListService(ctx).GetVideos(req)
	if err != nil {
		resp = &video.GetVideosResponse{StatusCode: unierr.GetVideoListFiled.ErrCode, StatusMsg: unierr.GetVideoListFiled.ErrMsg}
		return resp, err
	}

	resp = &video.GetVideosResponse{StatusCode: unierr.SuccessCode.ErrCode, StatusMsg: unierr.SuccessCode.ErrMsg, Videos: videoList}
	return resp, nil
}
