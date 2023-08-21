package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

var _ endpoint.Middleware = ServerMiddleware

// ServerMiddleware Kitex middleware: 截取 client 信息作为日志
func ServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)

		klog.Infof("Client address: %v\n", ri.From().Address())

		if err = next(ctx, req, resp); err != nil {
			return err
		}

		return nil
	}
}
