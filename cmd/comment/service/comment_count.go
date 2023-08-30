package service

import (
	"context"
	"minitok/cmd/comment/dal"
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
	commentCounts, err := dal.CountComments(req.VideoIdList, s.ctx)
	if err != nil {
		return nil, err
	}
	return commentCounts, nil
}
