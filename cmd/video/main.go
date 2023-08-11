package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"minitok/internal/conf"
	"minitok/internal/middleware"
	video "minitok/kitex_gen/video/videoservice"
	"net"
)

func initAll() {}

// TODO: 增加其它配置
func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.ETCDAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", conf.VideoServiceAddress)
	if err != nil {
		panic(err)
	}

	initAll()

	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.VideoServiceName}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
