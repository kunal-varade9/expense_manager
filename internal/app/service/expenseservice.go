package service

import (
	"context"
)

type ExpenseService struct {
	BaseService
}

func InitExpenseService(ctx context.Context) *ExpenseService {
	return &ExpenseService{
		*InitBaseService(ctx),
	}
}
