package service

import (
	"context"
	"minitok/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

func (s *CommentListService) CommentList(req *comment.ListRequest) ([]*comment.Comment, error) {

	return nil, nil
}
