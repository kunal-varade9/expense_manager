package repository

import (
	"context"
	"expense-manager/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	ctx context.Context
	db  *mongo.Database
}

func InitRepository(ctx context.Context) *Repository {
	instance := service.GetInstance()
	db := instance.GetConnection()

	return &Repository{
		ctx: ctx,
		db:  db,
	}
}
