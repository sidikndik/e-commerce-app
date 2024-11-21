package main

import (
	"e-commerce-app/infra"
	"e-commerce-app/router"

	"go.uber.org/zap"
)

func main() {
	ctx, err := infra.NewContext()
	if err != nil {
		ctx.Log.Panic("Error", zap.Error(err))
		return
	}

	router.SetupReouter(ctx)
}
