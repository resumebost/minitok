package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"minitok/cmd/favorite/dal"
	"minitok/cmd/favorite/rpc"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/oss"
	"minitok/internal/tracer"
	favorite "minitok/kitex_gen/favorite/favoriteservice"
	"net"
)

var favoriteConstants *constant.ServiceInfo

func initAll() {
	constant.InitConstant()
	favoriteConstants = &constant.AllConstants.FavoriteService

	dal.SetFavoriteDB()
	rpc.InitForFavorite()
	oss.InitOSS()
	tracer.InitJaeger(favoriteConstants.Name)
}

// TODO: 增加其它配置
func main() {
	initAll()

	r, err := etcd.NewEtcdRegistry([]string{constant.ETCDAddress})
	if err != nil {
		klog.Fatalf("ETCD registry initialization failed: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", favoriteConstants.Addr())
	if err != nil {
		klog.Fatalf("Unable to obtain TCP address: %v", err)
	}

	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: favoriteConstants.Name}),
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
