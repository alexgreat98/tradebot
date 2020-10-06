package providers

import (
	"context"
	"github.com/webdelo/tradebot/providers/bindings"
)

func Bind(ctx context.Context) {
	bindings.BindRepositories(ctx)
}
