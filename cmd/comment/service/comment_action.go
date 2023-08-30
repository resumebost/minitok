package service

import (
	"context"
	"minitok/kitex_gen/comment"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{
		ctx: ctx,
	}
}

func (s *CommentActionService) PostComment(req *comment.ActionRequest) (*comment.Comment, error) {

	return nil, nil
}

func (s *CommentActionService) DeleteComment(req *comment.ActionRequest) error {

	return nil
}
