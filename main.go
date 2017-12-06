package main

import (
	"context"

	"github.com/fnproject/fn/api/server"
	_ "github.com/treeder/fn-ext-example/logspam"
	// _ "github.com/treeder/fn-ext-example2/logspam2"
)

func main() {
	ctx := context.Background()
	funcServer := server.NewFromEnv(ctx)
	funcServer.AddExtensionByName("github.com/treeder/fn-ext-example/logspam")
	funcServer.Start(ctx)
}
