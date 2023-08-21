package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"minitok/cmd/video/dal"
	"minitok/cmd/video/rpc"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/oss"
	"minitok/internal/tracer"
	video "minitok/kitex_gen/video/videoservice"
	"net"
)

var videoConstants *constant.ServiceInfo

func initAll() {
	constant.InitConstant()
	videoConstants = &constant.AllConstants.VideoService

	dal.SetVideoDB()
	rpc.InitForVideo()
	oss.InitOSS()
	tracer.InitJaeger(videoConstants.Name)
}

func main() {
	initAll()

	r, err := etcd.NewEtcdRegistry([]string{constant.ETCDAddress})
	if err != nil {
		klog.Fatalf("ETCD registry initialization failed: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", videoConstants.Addr())
	if err != nil {
		klog.Fatalf("Unable to obtain TCP address: %v", err)
	}

	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: videoConstants.Name}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatalf("Kitex server can not run: %v", err)
	}
}
