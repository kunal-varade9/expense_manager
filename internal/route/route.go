package route

import (
	"context"
	"expense-manager/internal/app/handler"
	"expense-manager/internal/app/middleware"
	"expense-manager/internal/utils"

	"github.com/gin-gonic/gin"
)

type route struct {
	ctx     context.Context
	routers *gin.RouterGroup
}

func Routers(ctx context.Context, rGroup *gin.RouterGroup) {
	r := route{
		ctx:     ctx,
		routers: rGroup.Group("v1"),
	}
	utils.ExecuteMethods(&r)
}

func (r *route) UserRoutes() {
	userGroup := r.routers.Group("user")
	h := handler.InitHandler(r.ctx)
	userGroup.GET("/getuser", middleware.ErrorWrapper(h.GetAllExpenses))
}

func (r *route) ExpenseRoutes() {
	expenseGroup := r.routers.Group("expense")
	h := handler.InitHandler(r.ctx)
	expenseGroup.GET("/get", middleware.ErrorWrapper(h.GetAllExpenses))
}
