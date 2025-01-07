package main

import (
	"context"
	"log"

	"github.com/ndreyserg/gophermart/internal/app"
)

func main() {
	ctx := context.Background()
	app, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("init app failed %s", err.Error())
	}
	err = app.Run()
	if err != nil {
		log.Fatalf("run app failed %s", err.Error())
	}
}
