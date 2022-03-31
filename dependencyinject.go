package main

import (
	"context"

	"github.com/rfscheidt/goapi/http"
)

func dependencyinject() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "key", "any value in the context")

	api := http.NewAPI(ctx)

	api.SetupServer()
}
