package service

import (
	"context"
	"minitok/kitex_gen/video"
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
	//TODO
	return nil, nil
}
