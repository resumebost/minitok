package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"minitok/cmd/user/dal"
	"minitok/cmd/user/rpc"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/oss"
	"minitok/internal/tracer"
	user "minitok/kitex_gen/user/userservice"
	"net"
)

var userConstants *constant.ServiceInfo

func initAll() {
	constant.InitConstant()
	userConstants = &constant.AllConstants.UserService

	dal.SetUserDB()
	rpc.InitForUser()
	oss.InitOSS()
	tracer.InitJaeger(userConstants.Name)
}

// TODO: 增加其它配置
func main() {
	initAll()

	r, err := etcd.NewEtcdRegistry([]string{constant.ETCDAddress})
	if err != nil {
		klog.Fatalf("ETCD registry initialization failed: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", userConstants.Addr())
	if err != nil {
		klog.Fatalf("Unable to obtain TCP address: %v", err)
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: userConstants.Name}),
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
