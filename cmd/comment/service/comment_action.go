package service

import (
	"context"
	"minitok/cmd/comment/dal"
	"minitok/cmd/comment/rpc"
	"minitok/internal/jwt"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/user"
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
	videoId := req.VideoId
	commentContent := req.CommentText
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	userId := claims.ID

	createComment, err := dal.CreateComment(&dal.Comment{
		UserID:  userId,
		VideoID: videoId,
		Content: commentContent,
	}, s.ctx)
	if err != nil {
		return nil, err
	}

	author, err := rpc.GetUserInfo(s.ctx, &user.InfoRequest{
		UserId: createComment.UserID,
		Token:  req.Token,
	})
	if err != nil {
		return nil, err
	}

	res := &comment.Comment{
		Id:         int64(createComment.ID),
		User:       author.User,
		Content:    createComment.Content,
		CreateDate: createComment.CreatedAt.Format("01-02"), //MM-dd格式
	}

	return res, nil
}

func (s *CommentActionService) DeleteComment(req *comment.ActionRequest) error {
	commentId := req.CommentId
	err := dal.DeleteComment(commentId, s.ctx)
	return err
}
