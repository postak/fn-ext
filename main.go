package main

import (
	"context"

	"github.com/fnproject/fn/api/server"
	_ "./proxima_auth"
	_ "github.com/fnproject/fn/api/server/defaultexts"

)

func main() {
	ctx := context.Background()
	funcServer := server.NewFromEnv(ctx)
	funcServer.AddExtensionByName("github.com/postak/fn-ext/proxima_auth")
	funcServer.Start(ctx)
}
