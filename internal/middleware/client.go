package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

var _ endpoint.Middleware = ClientMiddleware

// ClientMiddleware Kitex middleware: 打印服务端地址, RPC 调用延迟和 IO 延迟
func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)

		klog.Infof("Server address: %v, RPC timeout: %v, readwrite timeout: %v\n", ri.To().Address(),
			ri.Config().RPCTimeout(), ri.Config().ConnectTimeout())

		if err = next(ctx, req, resp); err != nil {
			return err
		}

		return nil
	}
}
