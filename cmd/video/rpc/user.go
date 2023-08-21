package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"minitok/internal/conf"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/user"
	"minitok/kitex_gen/user/userservice"
	"time"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		conf.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}

	userClient = c
}

func GetUserInfo(ctx context.Context, req *user.InfoRequest) (*user.InfoResponse, error) {
	resp, err := userClient.Info(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
