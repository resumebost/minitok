package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

var _ endpoint.Middleware = ServerMiddleware

// ServerMiddleware TODO: 在请求进入 server 前进行拦截, 截取请求信息作为日志
func ServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		return nil
	}
}
