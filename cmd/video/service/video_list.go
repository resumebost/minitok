package service

import (
	"context"
	"minitok/kitex_gen/video"
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
	//TODO
	return nil, nil
}
