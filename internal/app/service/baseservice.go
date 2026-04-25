package service

import (
	"context"
	"expense-manager/internal/repository"
)

type BaseService struct {
	Repo repository.Repository
	ctx  context.Context
}

func InitBaseService(ctx context.Context) *BaseService {
	return &BaseService{
		ctx:  ctx,
		Repo: *repository.InitRepository(ctx),
	}
}
