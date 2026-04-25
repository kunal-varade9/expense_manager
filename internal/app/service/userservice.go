package service

import (
	"context"
)

type UserService struct {
	BaseService
}

func InitUserService(ctx context.Context) *UserService {
	return &UserService{
		*InitBaseService(ctx),
	}
}
