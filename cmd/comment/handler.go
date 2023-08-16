package main

import (
	"context"
	comment "minitok/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// Action implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Action(ctx context.Context, req *comment.ActionRequest) (resp *comment.ActionResponse, err error) {
	// TODO: Your code here...
	return
}

// List implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) List(ctx context.Context, req *comment.ListRequest) (resp *comment.ListResponse, err error) {
	// TODO: Your code here...
	return
}

// Count implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Count(ctx context.Context, req *comment.CountRequest) (resp *comment.CountResponse, err error) {
	// TODO: Your code here...
	return
}
