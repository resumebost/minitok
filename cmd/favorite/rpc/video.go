package rpc

import (
	"context"
	"fmt"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/unierr"
	"minitok/kitex_gen/video"
	"minitok/kitex_gen/video/videoservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
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

// 请求查询
// GetVideosInfo multiple get list of video info
func GetVideosInfo(ctx context.Context, videoIDs []int64) (map[int64]*video.Video, error) {
	req := &video.GetVideosRequest{
		VideoIds: videoIDs,
	}

	resp, err := videoClient.GetVideos(ctx, req)
	if err != nil {
		//TODO 
		fmt.Printf("Error in RPC call: %v\n", err)
		return nil, err
	}

  	// rpc响应异常处理
	  if resp.StatusCode != 0 {
		return nil, unierr.NewErrCore(resp.StatusCode, resp.StatusMsg)
	}

	videoMap := make(map[int64]*video.Video)
	for _, v := range resp.Videos {
		videoMap[v.Id] = v
	}

	return videoMap, nil
}
