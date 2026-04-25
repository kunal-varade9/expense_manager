package main

import (
	"context"
	"expense-manager/cmd/boot"
	"expense-manager/internal/route"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	settingEnvVariable()

	boot.Boot(ctx)

	app := gin.New()
	api := app.Group("api")
	route.Routers(ctx, api)

	app.Run()

	go terminate(cancel)
}

func terminate(cancel context.CancelFunc) {
	cancel()
}

func settingEnvVariable() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}
