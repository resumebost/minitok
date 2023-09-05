package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"minitok/cmd/api/rpc"
	"minitok/kitex_gen/comment"
	"testing"
)

func TestActionPost(t *testing.T) {
	doTestActionPost(t)
}

func TestActionDelete(t *testing.T) {
	doTestActionDelete(t)
}

func TestCommentList(t *testing.T) {
	doCommentList(t)
}

func TestCommentCount(t *testing.T) {
	doCommentCount(t)
}

func doTestActionPost(t *testing.T) {
	resp, err := rpc.CommentAction(ctx, &comment.ActionRequest{
		Token:       TOKEN,
		VideoId:     1,
		ActionType:  1,
		CommentText: "评论测试",
	})
	assert.NoError(t, err)
	fmt.Println(resp)
}

func doTestActionDelete(t *testing.T) {
	resp, err := rpc.CommentAction(ctx, &comment.ActionRequest{
		Token:      TOKEN,
		ActionType: 2,
		CommentId:  1,
	})
	assert.NoError(t, err)
	fmt.Println(resp)
}

func doCommentList(t *testing.T) {
	resp, err := rpc.CommentList(ctx, &comment.ListRequest{
		Token:   TOKEN,
		VideoId: 1,
	})
	assert.NoError(t, err)
	fmt.Println(resp)
}

func doCommentCount(t *testing.T) {
	resp, err := rpc.CommentCount(ctx, &comment.CountRequest{VideoIdList: []int64{1}})
	assert.NoError(t, err)
	fmt.Println(resp)
}
