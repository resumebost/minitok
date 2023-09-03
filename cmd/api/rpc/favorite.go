package rpc

import (
	"context"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/favorite"
	"minitok/kitex_gen/favorite/favoriteservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

var favoriteClient favoriteservice.Client
var favoriteConstants = &constant.AllConstants.FavoriteService

func initFavoriteRPC() {
	r, err := etcd.NewEtcdResolver([]string{constant.ETCDAddress})
	if err != nil {
		klog.Fatalf("ETCD resolver initialization failed: %v", err)
	}

	c, err := favoriteservice.NewClient(
		favoriteConstants.Name,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(opentracing.NewDefaultClientSuite()),
		client.WithResolver(r),
	)

	if err != nil {
		klog.Fatalf("Kitex client initialization failed: %v", err)
	}

	favoriteClient = c
}

func FavoriteAction(ctx context.Context, req *favorite.ActionRequest) (*favorite.ActionResponse, error) {
	resp, err := favoriteClient.Action(ctx, req)

	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteList(ctx context.Context, req *favorite.ListRequest) (*favorite.ListResponse, error) {
	resp, err := favoriteClient.List(ctx, req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteJudge(ctx context.Context, req *favorite.JudgeRequest) (*favorite.JudgeResponse, error) {
	resp, err := favoriteClient.Judge(ctx, req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteCount(ctx context.Context, req *favorite.CountRequest) (*favorite.CountResponse, error) {
	resp, err := favoriteClient.Count(ctx, req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteCountByUser(ctx context.Context, req *favorite.CountByUserRequest) (*favorite.CountByUserResponse, error) {
	resp, err := favoriteClient.CountByUser(ctx, req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
