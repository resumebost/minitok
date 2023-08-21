package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"minitok/cmd/comment/dal"
	"minitok/cmd/comment/rpc"
	"minitok/internal/constant"
	"minitok/internal/middleware"
	"minitok/internal/oss"
	"minitok/internal/tracer"
	comment "minitok/kitex_gen/comment/commentservice"
	"net"
)

var commentConstants *constant.ServiceInfo

func initAll() {
	constant.InitConstant()
	commentConstants = &constant.AllConstants.CommentService

	dal.SetCommentDB()
	rpc.InitForComment()
	oss.InitOSS()
	tracer.InitJaeger(commentConstants.Name)
}

// TODO: 增加其它配置
func main() {
	initAll()

	r, err := etcd.NewEtcdRegistry([]string{constant.ETCDAddress})
	if err != nil {
		klog.Fatalf("ETCD registry initialization failed: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", commentConstants.Addr())
	if err != nil {
		klog.Fatalf("Unable to obtain TCP address: %v", err)
	}

	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: commentConstants.Name}),
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
