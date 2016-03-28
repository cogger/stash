package stash

import "golang.org/x/net/context"

func Has(ctx context.Context, key string) bool {
	_, holder := c(ctx)
	holder.RLock()
	_, ok := holder.values[key]
	holder.RUnlock()
	return ok
}
