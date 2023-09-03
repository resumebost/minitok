package service

import (
    "context"
    "minitok/cmd/favorite/dal"
)

type CountByUserService struct {
    ctx context.Context
}

// 初始化ctx
func NewCountByUserService(ctx context.Context) *CountByUserService {
    return &CountByUserService{ctx: ctx}
}

// CountByUser 统计用户的收藏数量
func (s *CountByUserService) CountByUser(userID int64) (int64, error) {
    count, err := dal.CountFavoritesByUser(s.ctx, userID)
    if err != nil {
        return 0, err
    }
    return count, nil
}
