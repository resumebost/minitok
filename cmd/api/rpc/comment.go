package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/comment"
	"minitok/kitex_gen/comment/commentservice"
	"time"
)

var commentClient commentservice.Client
var commentConstants = &constant.AllConstants.CommentService

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{constant.ETCDAddress})
	if err != nil {
		klog.Fatalf("ETCD resolver initialization failed: %v", err)
	}

	c, err := commentservice.NewClient(
		commentConstants.Name,
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

	commentClient = c
}

func CommentAction(ctx context.Context, req *comment.ActionRequest) (*comment.ActionResponse, error) {
	resp, err := commentClient.Action(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *comment.ListRequest) (*comment.ListResponse, error) {
	resp, err := commentClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func CommentCount(ctx context.Context, req *comment.CountRequest) (*comment.CountResponse, error) {
	resp, err := commentClient.Count(ctx, req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
