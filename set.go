package stash

import "golang.org/x/net/context"

func Set(ctx context.Context, key string, value interface{}) context.Context {
	ctx, holder := c(ctx)
	holder.Lock()
	holder.values[key] = value
	holder.Unlock()
	return ctx
}
