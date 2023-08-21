package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"minitok/internal/conf"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/favorite"
	"minitok/kitex_gen/favorite/favoriteservice"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		conf.FavoriteServiceName,
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
	favoriteClient = c
}

func CountFavorite(ctx context.Context, req *favorite.CountRequest) (*favorite.CountResponse, error) {
	resp, err := favoriteClient.Count(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func JudgeFavorite(ctx context.Context, req *favorite.JudgeRequest) (*favorite.JudgeResponse, error) {
	resp, err := favoriteClient.Judge(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
