package main

import (
	"context"

	_ "github.com/fnproject/fn-ext-example/logspam"
	"github.com/fnproject/fn/api/server"
	_ "github.com/fnproject/fn/api/server/defaultexts"
)

func main() {
	ctx := context.Background()
	funcServer := server.NewFromEnv(ctx)

	funcServer.AddExtensionByName("github.com/fnproject/fn-ext-example/logspam")
	funcServer.Start(ctx)
}
