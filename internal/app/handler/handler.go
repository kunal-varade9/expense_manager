package handler

import (
	"context"
	"expense-manager/internal/app/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctx            context.Context
	userService    service.UserService
	expenseService service.ExpenseService
}

func InitHandler(ctx context.Context) *Handler {
	return &Handler{
		ctx:            ctx,
		userService:    *service.InitUserService(ctx),
		expenseService: *service.InitExpenseService(ctx),
	}
}

func (h *Handler) GetAllExpenses(c *gin.Context) error {
	var result any

	result, err := h.expenseService.BaseService.Repo.FetchExpenses(h.ctx)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"data": result,
	})
	return nil
}
