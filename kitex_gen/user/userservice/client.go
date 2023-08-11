// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "minitok/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Info(ctx context.Context, Req *user.InfoRequest, callOptions ...callopt.Option) (r *user.InfoResponse, err error)
	Register(ctx context.Context, Req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	Login(ctx context.Context, Req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Info(ctx context.Context, Req *user.InfoRequest, callOptions ...callopt.Option) (r *user.InfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Info(ctx, Req)
}

func (p *kUserServiceClient) Register(ctx context.Context, Req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kUserServiceClient) Login(ctx context.Context, Req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}
