package service

import (
	"context"
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

func (s *FeedService) Feed(req *video.FeedRequest) ([]*video.Video, int64, error) {
	//TODO
	return nil, 0, nil
}
