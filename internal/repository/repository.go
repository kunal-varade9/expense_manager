package repository

import (
	"context"
	"expense-manager/internal/app/service/mongoservice"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	ctx context.Context
	db  *mongo.Database
}

func InitRepository(ctx context.Context) *Repository {
	instance := mongoservice.GetInstance()
	db := instance.GetConnection()

	return &Repository{
		ctx: ctx,
		db:  db,
	}
}

func (r *Repository) FetchUsers(ctx context.Context) (any, error) {
	user_collection := r.db.Collection("users")
	cursor, err := user_collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var result any
	err = cursor.All(ctx, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FetchExpenses(ctx context.Context) (any, error) {
	fetch_expense := r.db.Collection("expenses")
	cursor, err := fetch_expense.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var result []bson.M
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
