package stash

import "golang.org/x/net/context"

func Remove(ctx context.Context, key string) context.Context {
	if !Has(ctx, key) {
		return ctx
	}

	ctx, holder := c(ctx)
	holder.Lock()
	delete(holder.values, key)
	holder.Unlock()

	return ctx
}
