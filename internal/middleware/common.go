package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

// 防止编译器报定义未使用的错误
var _ endpoint.Middleware = CommonMiddleware

// CommonMiddleware TODO: 将 RPC 调用信息截获作为日志
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		return nil
	}
}
