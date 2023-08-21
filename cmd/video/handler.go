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
	// TODO: Your code here...
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	//println(ctx.Value("id"))

	resp = nil

	err = service.NewUploadVideoService(ctx).PublishVideo(req)
	if err != nil {
		resp = &video.PublishActionResponse{StatusCode: unierr.VideoPublishFiled.ErrCode, StatusMsg: unierr.VideoPublishFiled.ErrMsg}
		return resp, err
	}

	resp = &video.PublishActionResponse{StatusCode: unierr.SuccessCode.ErrCode, StatusMsg: unierr.SuccessCode.ErrMsg}
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
