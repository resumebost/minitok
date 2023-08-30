package test

import (
	"context"
	"fmt"
	"testing"

	"minitok/cmd/api/rpc"
	"minitok/internal/constant"
	"minitok/internal/jwt"
	"minitok/kitex_gen/favorite"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background() // 定义上下文
var token, _ = jwt.GenToken(1, "xiayi")
func TestMain(m *testing.M) {
	constant.InitConstant() // 初始化配置
	rpc.InitRPC()           // 初始化rpc服务
	m.Run()
}

// go test -v ./test/favorite_test.go -run TestAction
func TestAction(t *testing.T) {
	doAction(ctx, t)
}
// go test -v ./test/favorite_test.go -run TestList
func TestList(t *testing.T) {
	doList(ctx,t)
}

// go test -v ./test/favorite_test.go -run TestJudge
func TestJudge(t *testing.T) {
	doJudge(ctx, t)
}

// go test -v ./test/favorite_test.go -run TestCount
func TestCount(t *testing.T) {
	doCount(ctx, t)
}

func doAction(ctx context.Context, t *testing.T) {
	resp, err := rpc.FavoriteAction(ctx, &favorite.ActionRequest{
		Token:      token,
		VideoId:    10,
		ActionType: 1,
	})
	assert.NoError(t, err)
	fmt.Printf("rpc服务响应：%v", resp)
}

func doList(ctx context.Context, t *testing.T) {
	resp, err := rpc.FavoriteList(ctx, &favorite.ListRequest{
		UserId: 1,
		Token:  token,
	})
	assert.NoError(t, err)
	fmt.Printf("rpc服务响应：%v", resp)
}

func doJudge(ctx context.Context, t *testing.T) {
	resp, err := rpc.FavoriteJudge(ctx, &favorite.JudgeRequest{
		Token:       token,
		VideoIdList: []int64{1, 2, 3, 6},
	})
	assert.NoError(t, err)
	fmt.Printf("rpc服务响应：%v", resp)
}

func doCount(ctx context.Context, t *testing.T) {
	resp, err := rpc.FavoriteCount(ctx, &favorite.CountRequest{
		VideoIdList: []int64{1, 2, 10},
	})
	assert.NoError(t, err)
	fmt.Printf("rpc服务响应：%v", resp)
}