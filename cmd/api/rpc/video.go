package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"minitok/internal/conf"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/video"
	"minitok/kitex_gen/video/videoservice"
	"time"
)

var videoClient videoservice.Client

// TODO: 增加更多配置
func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.ETCDAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		conf.VideoServiceName,
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

	videoClient = c
}

func VideoPublishAction(ctx context.Context, req *video.PublishActionRequest) (*video.PublishActionResponse, error) {
	resp, err := videoClient.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}
	return resp, nil
}
