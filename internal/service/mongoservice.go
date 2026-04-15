package service

import (
	"context"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	client *mongo.Client
	ctx    context.Context
}

var (
	instance *MongoService
	once     sync.Once
)

func Connect(ctx context.Context) *MongoService {

	once.Do(func() {

		uri := viper.GetString("MONGO_URI")
		opts := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			panic(err)
		}
		instance = &MongoService{
			client: client,
			ctx:    ctx,
		}
	})

	return instance
}

func GetInstance() *MongoService {
	return instance
}

func (ms *MongoService) GetConnection() *mongo.Database {
	db := ms.client.Database(viper.GetString("DATABASE_NAME"))
	if db == nil {
		panic("There is Connection Error")
	}
	return db
}

func (ms *MongoService) Close() error {
	err := ms.client.Disconnect(ms.ctx)
	if err != nil {
		return err
	}
	return nil
}
