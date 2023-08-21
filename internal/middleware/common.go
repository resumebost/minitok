package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

// 防止编译器报定义未使用的错误
var _ endpoint.Middleware = CommonMiddleware

// CommonMiddleware Kitex middleware: 打印 RPC 信息以及请求响应信息
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)

		klog.Infof("Real request: %+v\n", req)
		klog.Infof("Remote service name: %s, remote method: %s\n", ri.To().ServiceName(), ri.To().Method())

		if err = next(ctx, req, resp); err != nil {
			return err
		}

		klog.Infof("Real response: %+v\n", resp)
		return nil
	}
}
