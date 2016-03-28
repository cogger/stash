package stash

import "golang.org/x/net/context"

func Get(ctx context.Context, key string, load ...func() interface{}) interface{} {
	_, holder := c(ctx)
	holder.RLock()
	value, ok := holder.values[key]
	holder.RUnlock()

	if !ok && len(load) > 0 {
		value = load[0]()
		Set(ctx, key, value)
	}
	return value
}
