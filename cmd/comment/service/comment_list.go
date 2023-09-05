package service

import (
	"context"
	"minitok/cmd/comment/dal"
	"minitok/cmd/comment/rpc"
	"minitok/internal/unierr"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/user"
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
	videoId := req.VideoId

	comments, err := dal.CommentList(videoId, s.ctx)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return []*comment.Comment{}, nil
	}

	res := make([]*comment.Comment, len(comments))

	for i, c := range comments {
		author, err := rpc.GetUserInfo(s.ctx, &user.InfoRequest{
			UserId: c.UserID,
			Token:  req.Token,
		})
		if err != nil {
			return nil, unierr.NewErrCore(
				author.StatusCode,
				author.StatusMsg)
		}

		res[i] = &comment.Comment{
			Id:   int64(c.ID),
			User: author.User,
			//User:       &user.User{},
			Content:    c.Content,
			CreateDate: c.CreatedAt.Format("01-02"),
		}
	}
	return res, nil
}
