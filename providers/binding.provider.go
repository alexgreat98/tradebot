package providers

import (
	"context"
	"github.com/webdelo/tradebot/database"
	"github.com/webdelo/tradebot/providers/bindings"
)

// Boot all needed dependencies
func Boot(ctx *context.Context) error {
	sqlite, err := database.Sqlite()
	if err != nil {
		return err
	}

	bindings.BindRepositories(ctx, sqlite)
	return nil
}
