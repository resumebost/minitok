package test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"minitok/cmd/api/rpc"
	"minitok/internal/constant"
	"minitok/internal/jwt"
	"minitok/kitex_gen/user"
	"testing"
)

var ctxUser = context.Background() // 定义上下文
var tokenUser, _ = jwt.GenToken(4139418956861440, "xiayi")

func TestUser(t *testing.T) {
	constant.InitConstant() // 初始化配置
	rpc.InitRPC()           // 初始化rpc服务
}

func TestRegister(t *testing.T) {
	doRegister(ctxUser, t)
}

func TestLogin(t *testing.T) {
	doLogin(ctxUser, t)
}

func TestInfo(t *testing.T) {
	doInfo(ctxUser, t)
}

func doRegister(ctxUser context.Context, t *testing.T) {
	req := &user.RegisterRequest{
		Username: "xiayi3",
		Password: "123456",
	}
	resp, err := rpc.Register(ctxUser, req)
	assert.NoError(t, err)
	fmt.Printf("rpc服务响应：%v", resp)
}

func doLogin(ctxUser context.Context, t *testing.T) {
	req := &user.LoginRequest{
		Username: "xiayi",
		Password: "123456",
	}
	resp, err := rpc.Login(ctxUser, req)
	assert.NoError(t, err)
	fmt.Printf("rpc服务响应：%v", resp)
}

func doInfo(ctxUser context.Context, t *testing.T) {
	req := &user.InfoRequest{
		UserId: 4139418956861440,
		Token:  tokenUser,
	}
	resp, err := rpc.Info(ctxUser, req)
	assert.NoError(t, err)
	fmt.Printf("rpc服务响应：%v", resp)
}
