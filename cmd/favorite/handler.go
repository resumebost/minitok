package main

import (
	"context"
	favorite "minitok/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// Action implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Action(ctx context.Context, req *favorite.ActionRequest) (resp *favorite.ActionResponse, err error) {
	// TODO: Your code here...
	return
}

// List implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) List(ctx context.Context, req *favorite.ListRequest) (resp *favorite.ListResponse, err error) {
	// TODO: Your code here...
	return
}

// Judge implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Judge(ctx context.Context, req *favorite.JudgeRequest) (resp *favorite.JudgeResponse, err error) {
	// TODO: Your code here...
	return
}

// Count implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Count(ctx context.Context, req *favorite.CountRequest) (resp *favorite.CountResponse, err error) {
	// TODO: Your code here...
	return
}
