package main

import (
	"context"

	"github.com/fnproject/fn/api/server"
	"github.com/treeder/fn-plugin-example/logspam"
)

func main() {
	ctx := context.Background()

	funcServer := server.NewFromEnv(ctx)
	// Setup your custom extensions, listeners, etc here
	spammer := &logspam.LogSpam{}
	funcServer.AddCallListener(spammer)

	funcServer.Start(ctx)
}
