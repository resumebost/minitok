package main

import (
	"context"
	"minitok/cmd/comment/service"
	"minitok/internal/unierr"
	"minitok/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// Action implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Action(ctx context.Context, req *comment.ActionRequest) (resp *comment.ActionResponse, err error) {
	resp = nil
	var commentRes *comment.Comment = nil

	if req.ActionType == 1 {
		commentRes, err = service.NewCommentActionService(ctx).PostComment(req)
		if err != nil {
			resp = &comment.ActionResponse{
				StatusCode: unierr.CommentPostFiled.ErrCode,
				StatusMsg:  unierr.CommentPostFiled.ErrMsg,
			}
			return resp, err
		}
	} else if req.ActionType == 2 {
		err = service.NewCommentActionService(ctx).DeleteComment(req)
		if err != nil {
			resp = &comment.ActionResponse{
				StatusCode: unierr.CommentDeleteFiled.ErrCode,
				StatusMsg:  unierr.CommentDeleteFiled.ErrMsg,
			}
			return resp, err
		}
	}

	resp = &comment.ActionResponse{
		StatusCode: unierr.SuccessCode.ErrCode,
		StatusMsg:  unierr.SuccessCode.ErrMsg,
		Comment:    commentRes,
	}
	return resp, nil
}

// List implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) List(ctx context.Context, req *comment.ListRequest) (resp *comment.ListResponse, err error) {
	resp = nil

	commentList, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp = &comment.ListResponse{
			StatusCode: unierr.CommentNotFound.ErrCode,
			StatusMsg:  unierr.CommentNotFound.ErrMsg,
		}
	}

	resp = &comment.ListResponse{
		StatusCode:  unierr.SuccessCode.ErrCode,
		StatusMsg:   unierr.SuccessCode.ErrMsg,
		CommentList: commentList,
	}

	return resp, nil
}

// Count implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Count(ctx context.Context, req *comment.CountRequest) (resp *comment.CountResponse, err error) {
	resp = nil

	countList, err := service.NewCommentCountService(ctx).CommentCount(req)
	if err != nil {
		resp = &comment.CountResponse{
			StatusCode: unierr.InternalError.ErrCode,
			StatusMsg:  unierr.InternalError.ErrMsg,
		}
	}

	resp = &comment.CountResponse{
		StatusCode:       unierr.SuccessCode.ErrCode,
		StatusMsg:        unierr.SuccessCode.ErrMsg,
		CommentCountList: countList,
	}

	return resp, nil
}
