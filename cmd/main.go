package main

import (
	"context"
	"urlshorter/cmd/app"
	"urlshorter/config"
)

func main() {
	ctx := context.Background()
	config := config.NewConfig()
	app := app.NewServer(config)
	app.Start(ctx)
}
