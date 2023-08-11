package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

var _ endpoint.Middleware = ClientMiddleware

// ClientMiddleware TODO: 将 timeout 作为日志
func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		return nil
	}
}
