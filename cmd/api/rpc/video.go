package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/video"
	"minitok/kitex_gen/video/videoservice"
	"time"
)

var videoClient videoservice.Client
var videoConstants = &constant.AllConstants.VideoService

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constant.ETCDAddress})
	if err != nil {
		klog.Fatalf("ETCD resolver initialization failed: %v", err)
	}

	c, err := videoservice.NewClient(
		videoConstants.Name,
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

	videoClient = c
}

func VideoPublishAction(ctx context.Context, req *video.PublishActionRequest) (*video.PublishActionResponse, error) {
	resp, err := videoClient.PublishAction(ctx, req, callopt.WithRPCTimeout(180*time.Second))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		//fmt.Println("rpc return resp.StatusCode!=0")
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func VideoPublishList(ctx context.Context, req *video.PublishListRequest) (*video.PublishListResponse, error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func GetVideoList(ctx context.Context, req *video.GetVideosRequest) (*video.GetVideosResponse, error) {
	resp, err := videoClient.GetVideos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}

func Feed(ctx context.Context, req *video.FeedRequest) (*video.FeedResponse, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
