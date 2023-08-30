package service

import (
	"context"
	"minitok/kitex_gen/comment"
)

type CommentCountService struct {
	ctx context.Context
}

func NewCommentCountService(ctx context.Context) *CommentCountService {
	return &CommentCountService{
		ctx: ctx,
	}
}

func (s *CommentCountService) CommentCount(req *comment.CountRequest) ([]int64, error) {

	return nil, nil
}
