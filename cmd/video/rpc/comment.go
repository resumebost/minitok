package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"minitok/internal/conf"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/comment/commentservice"
	"time"
)

var commentClient commentservice.Client

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		conf.CommentServiceName,
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
	commentClient = c
}

func CountComment(ctx context.Context, req *comment.CountRequest) (*comment.CountResponse, error) {
	resp, err := commentClient.Count(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
