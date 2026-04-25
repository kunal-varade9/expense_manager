package boot

import (
	"context"
	"expense-manager/internal/app/service/mongoservice"
)

func Boot(ctx context.Context) bool {
	// will be responsible for all connection to be stablise
	mongoservice.Connect(ctx)

	return true
}
