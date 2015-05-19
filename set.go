package stash

import "golang.org/x/net/context"

func Set(ctx context.Context, key string, value interface{}) context.Context {
	values := c(ctx)
	values[key] = value

	return context.WithValue(ctx, &stashkey, &values)
}
